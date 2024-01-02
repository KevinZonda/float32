import {observer} from "mobx-react-lite";
import ReqStore from "./Store/ReqStore.ts";
import {Button} from "tdesign-react";

export const RelatedQuestion = observer(() => {
  if (ReqStore.relatedList && ReqStore.relatedList.length === 0) {
    return <></>
  }
  return (
    <div style={{textAlign: 'left'}}>
      <h3>❓相关问题</h3>
      {
        ReqStore.relatedList.map((v) => {
          return <Button style={{
            textAlign: 'left', justifyContent: 'left', width: '100%', marginBottom: '8px'
          }} theme="default" variant="outline">{v}</Button>
        })
      }
    </div>

  )
})