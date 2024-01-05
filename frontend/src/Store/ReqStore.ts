import {makeAutoObservable, runInAction} from "mobx";
import {BaseStore} from "./BaseStore.ts";

const baseAPI = 'https://api.float32.app/query'
const historyAPI = 'https://api.float32.app/history?id='
const continueAPI = 'https://api.float32.app/continue'

class reqStore {
  public shareLink(shareId: string) {
    if (!shareId || shareId === '') {
      return 'https://float32.app'
    }
    return 'https://float32.app/search?id=' + shareId
  }

  public constructor() {
    makeAutoObservable(this)
  }

  public isRainbow = false

  public evidenceList: Array<Evidence> = []
  public relatedList: Array<string> = []

  //region isLoading
  public get isLoading() {
    return this._isLoading
  }

  public set isLoading(v: boolean) {
    this._isLoading = v
  }

  public _isLoading: boolean = false
  //endregion

  public isFailed: boolean = false
  public shareId = ''
  public question = ''
  public warning = ''
  public currentAns: string = ''
  private _currentHistory = ''
  public prevQA: PrevAnsItem[] = []
  private _parentId = ''

  public get autoPrevOA() {
    if (this.isLoading) return this.prevQA
    if (this.prevQA.length < 1) return []
    return this.prevQA.slice(0, this.prevQA.length - 1)
  }

  private resetCore() {
    this.isFailed = false
    this.currentAns = ''
    this.evidenceList = []
  }

  public async queryHistory(id: string) {
    if (this.isLoading) return
    if (id === this._currentHistory) return
    this._currentHistory = id
    this.shareId = id
    this._parentId = id
    this.prevQA = []

    this.resetCore()
    this.isLoading = true
    await fetch(historyAPI + id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
    }).then(async (res) => {
      // decode res to json
      res.json().then((json) => {
        runInAction(() => {
          this.isLoading = false
          this.isFailed = false
          this.currentAns = json.answer ?? ''
          BaseStore.question = this.question = json.question ?? ''
          this.evidenceList = json.evidence ?? []
          this.relatedList = json.related ?? []
          this.pushQA(id, this.question, this.currentAns, this.evidenceList)
        })
      })

    }).catch((e) => {
      runInAction(() => {
        this.isFailed = true
        this.isLoading = false
        this.currentAns = 'Error: ' + e
      })
      return
    })
  }

  private async afterResponse(question: string, resp: Promise<Response>) {
    resp.then(async (res) => {
      let buf = ''
      let hasDoneMeta = false
      const reader = res.body!.pipeThrough(new TextDecoderStream()).getReader();
      /*eslint no-constant-condition: ["error", { "checkLoops": false }]*/
      while (true) {
        const {done, value} = await reader.read();
        if (value !== undefined) {
          if (hasDoneMeta) {
            this.currentAns = this.currentAns + value
            continue
          }
          buf = buf + value
          const idx = buf.indexOf('\r\n')
          if (idx !== -1) {
            const meta = buf.slice(0, idx)
            console.log(meta)
            const metaObj: AnsMetaInfo = JSON.parse(meta)
            this.evidenceList = metaObj.evidences
            this.currentAns = buf.slice(idx + 2)
            this.shareId = metaObj.id
            this._parentId = metaObj.id
            this.relatedList = metaObj.related ?? []
            window.history.replaceState(null, '', '/search?id=' + metaObj.id)
            hasDoneMeta = true
            continue
          }
        }
        this.pushQA(this.shareId, question, this.currentAns, this.evidenceList)

        if (done) break;
      }
      this.isLoading = false
    }).catch((r) => {
      this.isFailed = true
      this.isLoading = false
      this.currentAns = 'Error: ' + r
      return
    })
  }

  public async queryQuestion(question: string, lang: string, field: string, spec: string) {
    if (this.isLoading) return

    this.resetCore()
    this.isLoading = true
    this.shareId = ''
    this._parentId = ''
    this.prevQA = []

    const fresp = fetch(baseAPI, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        'question': question,
        'language': lang,
        'prog_lang': spec,
        'field': field,
      })
    })
    await this.afterResponse(question, fresp)
  }

  private pushQA(shareId: string, question: string, answer: string, evidence: Evidence[]) {
    this.prevQA.push({
      question: question,
      answer: answer,
      evidence: evidence,
      shareId: shareId
    })
  }


  public async continuousQuery(question: string, lang: string, field: string, spec: string) {
    if (this.isLoading) return
    this.resetCore()
    this.isLoading = true
    const fresp = fetch(continueAPI, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        'question': question,
        'language': lang,
        'prog_lang': spec,
        'field': field,
        'parent_id': this._parentId,
      })
    })
    await this.afterResponse(question, fresp)
  }
}

export interface AnsMetaInfo {
  evidences: Evidence[]
  id: string
  related: string[]
}

export interface Evidence {
  url: string
  title: string
  description: string
}

export interface PrevAnsItem {
  question: string
  answer: string
  evidence: Evidence[]
  shareId: string
}

export default new reqStore()
