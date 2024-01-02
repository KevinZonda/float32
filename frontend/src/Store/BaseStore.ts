import {fieldsOpt, langOpt} from "./const.tsx";

class baseStore {
  public constructor() {
  }

  public lang = langOpt[0].value
  public field = fieldsOpt[0].value
  public fieldSpec = this.field.options[0].value
  public question = ''
}

export const BaseStore = new baseStore()