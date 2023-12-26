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
        <Col span={ReqStore.evidenceList && ReqStore.evidenceList.length > 0 ? 4 : 0}>
          <List>
            {ReqStore.evidenceList && ReqStore.evidenceList.map((item) => (
              <ListItem>
                <Link href={item.url}>{item.title}</Link>
              </ListItem>
            ))}
          </List>
        </Col>
      </Row>
    </>)
})