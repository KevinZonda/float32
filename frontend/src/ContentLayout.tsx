import {Col, Link, List, Row} from "tdesign-react";
import {Content} from "./Content.tsx";
import {observer} from "mobx-react-lite";
import ReqStore from "./Store/ReqStore.ts";
import ListItem from "tdesign-react/es/list/ListItem";
import {useEffect, useState} from "react";

export const ContentLayout = observer(() => {
  const [isMobile, setIsMobile] = useState(window.innerWidth < 720)

//choose the screen size
  const handleResize = () => {
    if (window.innerWidth < 720) {
      setIsMobile(true)
    } else {
      setIsMobile(false)
    }
  }

// create an event listener
  useEffect(() => {
    window.addEventListener("resize", handleResize)
  })

  if (isMobile) {
    return (
      <>
        <Content/>
        <div style={{height: '24px'}}/>
        <Evidence/>
      </>
    )
  }

  return (
    <>
      <Row>
        <Col span={ReqStore.evidenceList && ReqStore.evidenceList.length > 0 ? 8 : 12}>
          <Content/>
        </Col>
        <Col style={{textAlign: 'left', paddingLeft: '24px'}}
             span={ReqStore.evidenceList && ReqStore.evidenceList.length > 0 ? 4 : 0}>
          <Evidence/>
        </Col>
      </Row>
    </>)
})

export const Evidence = observer(() => {
  if (!ReqStore.evidenceList || ReqStore.evidenceList.length === 0) {
    return <></>
  }
  return (
    <>
      <h3 style={{
        paddingBottom: '16px',
        marginBlock: 0,
        marginBlockStart: 0,
        marginBlockEnd: 0,
        textAlign: 'left'
      }}>{'📖 References'}</h3>
      <List style={{textAlign: 'left'}}>
        {ReqStore.evidenceList && ReqStore.evidenceList.map((item, idx) => (
          <ListItem style={{paddingTop: 0, paddingLeft: 0}}>
            <LinkBox title={item.title} url={item.url} idx={idx} description={item.description}/>
          </ListItem>
        ))}
      </List>
    </>
  )
})

interface LinkBoxProps {
  title: string
  url: string
  idx?: number
  description: string
}

const LinkBox = (props: LinkBoxProps) => {
  const uri = new URL(props.url).host
  return (
    <div style={{wordWrap: 'break-word'}}>
      <div>
        <img style={{verticalAlign: 'middle'}} src={'https://s2.googleusercontent.com/s2/favicons?domain=' + uri}></img>
        <span style={{color: 'grey', paddingLeft: '4px', verticalAlign: 'middle'}}>{uri}</span>
      </div>

      <Link theme="default" hover={'color'} href={props.url}>
        {props.title}
      </Link>
      <p style={{color: 'grey', margin: 0}}>
        {props.description}
      </p>
    </div>
  )
}