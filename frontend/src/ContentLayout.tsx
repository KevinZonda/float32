import {Col, Link, List, Row} from "tdesign-react";
import {Content} from "./Content.tsx";
import {observer} from "mobx-react-lite";
import ReqStore from "./Store/ReqStore.ts";
import ListItem from "tdesign-react/es/list/ListItem";

export const ContentLayout = observer(() => {
  return (
    <>
      <Row>
        <Col span={ReqStore.evidenceList && ReqStore.evidenceList.length > 0 ? 8 : 12}>
          <Content/>
        </Col>
        <Col style={{textAlign: 'left', paddingLeft: '24px'}} span={ReqStore.evidenceList && ReqStore.evidenceList.length > 0 ? 4 : 0}>
          <h3 style={{
            paddingBottom: '16px',
            marginBlock: 0,
            marginBlockStart: 0,
            marginBlockEnd: 0,
            textAlign: 'left'
          }}>{'📖 References'}</h3>
          <List>
            {ReqStore.evidenceList && ReqStore.evidenceList.map((item, idx) => (
              <ListItem style={{paddingTop: 0, paddingLeft: 0}}>
                <Link theme="default" hover="underline" href={item.url}>{'[' + (idx + 1) + '] ' + item.title}</Link>
              </ListItem>
            ))}
          </List>
        </Col>
      </Row>
    </>)
})