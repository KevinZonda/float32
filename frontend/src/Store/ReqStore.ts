import {makeAutoObservable} from "mobx";
const baseAPI = 'https://api.float32.ai/query'
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

  public async queryQuestion(question : string, lang: string, progLang: string) {
    if (this.isLoading) return

    this.isLoading = true
    this.currentAns = ''

    const res = await fetch(baseAPI, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        'question' : question,
        'language' : lang,
        'prog_lang' : progLang
      })
    })
    const reader = res.body!.pipeThrough(new TextDecoderStream()).getReader();
    /*eslint no-constant-condition: ["error", { "checkLoops": false }]*/
    while (true) {
      const { done, value } = await reader.read();
      console.log(value)
      if (value !== undefined) {
        this.currentAns = this.currentAns + value
      }
      if (done) break;
    }
    this.isLoading = false
  }

}

export default new reqStore()
