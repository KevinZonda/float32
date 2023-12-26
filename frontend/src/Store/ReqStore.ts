import {makeAutoObservable} from "mobx";

const baseAPI = 'https://api.float32.app/query'

class reqStore {

  private _warning = ''
  public get warning(): string {
    return this._warning
  }

  public set warning(value: string) {
    this._warning = value
  }
  public constructor() {
    makeAutoObservable(this)
  }

  private _currentAns: string = ''
  public get currentAns(): string {
    return this._currentAns
  }

  public set currentAns(value: string) {
    this._currentAns = value
  }

  private _evidenceList: Array<Evidence> = []
  public get evidenceList(): Array<Evidence> {
    return this._evidenceList
  }

  public set evidenceList(value: Array<Evidence>) {
    this._evidenceList = value
  }

  private _isLoading: boolean = false

  public get isLoading(): boolean {
    return this._isLoading
  }

  public set isLoading(value: boolean) {
    this._isLoading = value
  }

  private _isFailed: boolean = false
  public get isFailed(): boolean {
    return this._isFailed
  }

  public set isFailed(value: boolean) {
    this._isFailed = value
  }

  public async queryQuestion(question: string, lang: string, field : string, progLang: string) {
    if (this.isLoading) return

    this.isLoading = true
    this.isFailed = false
    this.currentAns = ''
    this.evidenceList = []

    await fetch(baseAPI, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        'question': question,
        'language': lang,
        'prog_lang': progLang,
        'field': field,
      })
    }).then(async (res) => {
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
            const metaObj : AnsMetaInfo = JSON.parse(meta)
            this.evidenceList = metaObj.evidences
            this.currentAns = buf.slice(idx + 2)
            hasDoneMeta = true
            continue
          }
        }
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
}

export interface AnsMetaInfo {
  evidences: Evidence[]
}

export  interface Evidence {
  url: string
  title: string
  description: string
}

export default new reqStore()
