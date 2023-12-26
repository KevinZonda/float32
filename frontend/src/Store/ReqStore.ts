import {makeAutoObservable} from "mobx";
const baseAPI = 'https://api.float32.app/query'
class reqStore {
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

  public async queryQuestion(question : string, lang: string, progLang: string) {
    if (this.isLoading) return

    this.isLoading = true
    this.isFailed = false
    this.currentAns = ''

    await fetch(baseAPI, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        'question' : question,
        'language' : lang,
        'prog_lang' : progLang
      })
    }).then(async (res) => {
      const reader = res.body!.pipeThrough(new TextDecoderStream()).getReader();
      /*eslint no-constant-condition: ["error", { "checkLoops": false }]*/
      while (true) {
        const { done, value } = await reader.read();
        if (value !== undefined) {
          this.currentAns = this.currentAns + value
        }
        if (done) break;

        this.isLoading = false
      }
    }).catch((r) => {
      this.isFailed = true
      this.isLoading = false
      this.currentAns = 'Error: ' + r
      return
    })
  }
}

export default new reqStore()
