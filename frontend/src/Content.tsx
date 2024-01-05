import {Button, NotificationPlugin, Popup, Skeleton} from "tdesign-react";
import ReqStore from "./Store/ReqStore.ts";
import {observer} from "mobx-react-lite";
import {FileCopyIcon, RefreshIcon, ShareIcon} from "tdesign-icons-react";
import rehypeKatex from 'rehype-katex'
import remarkMath from 'remark-math'
import remarkGfm from 'remark-gfm'
import MarkdownPreview from '@uiw/react-markdown-preview';

import './Warning.css'
import 'katex/dist/katex.min.css'
import {BaseStore} from "./Store/BaseStore.ts";
import React, {useEffect, useState} from "react";

interface operAnswerBtnProps {
  onClick: () => void
  icon: React.ReactElement
  hoverContent: string
}

const OperAnswerBtn = (props: operAnswerBtnProps) => {
  const [isMobile, setIsMobile] = useState(window.innerWidth < 720)
  useEffect(() => {
    window.addEventListener("resize", () => setIsMobile(window.innerWidth < 720))
  })
  if (isMobile) {
    return (
      <Button icon={props.icon} style={{marginRight: '8px'}} shape="square" variant="outline"
              onClick={props.onClick}>
      </Button>
    )
  }
  return (
    <Popup trigger="hover" content={props.hoverContent}>
      <Button icon={props.icon} style={{marginRight: '8px'}} shape="square" variant="outline"
              onClick={props.onClick}>
      </Button>
    </Popup>
  )
}

function notifySuccess(title: string, msg: string) {
  NotificationPlugin.success({
    title: title,
    content: msg,
    offset: [-10, 10],
    placement: 'top-right',
    duration: 1000,
    closeBtn: true,
  });
}

const Warning = observer(() => {
  return (
    ReqStore.warning === '' ? <></> :
      <>
        <div className="warning" style={{marginBottom: '16px'}}>
          <div style={{backgroundColor: '#fcfbfa', height: '100%', padding: '10px'}}>
            <h2 style={{
              paddingBottom: '16px',
              marginBlock: 0,
              marginBlockStart: 0,
              marginBlockEnd: 0,
              padding: 0,
              textAlign: 'left'
            }}>⚠️ 警告</h2>
            <p style={{textAlign: 'left', paddingBottom: 0, marginBlockEnd: 0, marginBottom: 0}}>
              {ReqStore.warning}
            </p>
          </div>


        </div>
      </>
  )
})

export const Content = observer((
  {loading, failed, question, text}: {loading : boolean, failed : boolean, question: string,  text: string }) => {
  if (!loading && text === '') {
    return <>
      <Warning/>
    </>
  }
  return (
    <>
      <Warning/>
      {
        question && question !== '' &&
          <>
              <h3 style={{
                paddingBottom: '16px',
                marginBlock: 0,
                marginBlockStart: 0,
                marginBlockEnd: 0,
                textAlign: 'left'
              }}>{'🤔 Question'}</h3>
              <p style={{textAlign: 'left', fontSize: '16px', paddingBottom: 0, marginBlockStart: 0}}>
                {question}
              </p>
          </>
      }
      <h3 style={{
        marginBlock: 0,
        marginBlockStart: 0,
        marginBlockEnd: 0,
        textAlign: 'left',
        marginBottom: '16px'
      }}>{failed ? '⚠️ Error' : '🔍 Answer'}</h3>
      <div style={{textAlign: 'left'}}>
        <MarkdownPreview
          remarkPlugins={[remarkGfm, remarkMath]} rehypePlugins={[rehypeKatex]}
          source={regularizeMarkdown(text)}
          wrapperElement={{
            "data-color-mode": "light"
          }}
        />
      </div>

      {loading ?
        <Skeleton animation={'flashed'} theme={'paragraph'} style={{paddingTop: '16px', paddingBottom: '16px'}}>
          <p>LOAD</p>
        </Skeleton> :
        <div style={{textAlign: 'left', paddingTop: '16px', paddingBottom: '16px'}}>
          <OperAnswerBtn icon={<FileCopyIcon/>}
                         hoverContent="复制答案"
                         onClick={() => {
                           navigator.clipboard.writeText(regularizeMarkdown(text))
                           notifySuccess('内容', '复制成功');
                         }}
          />
          <OperAnswerBtn icon={<ShareIcon/>}
                         hoverContent="复制分享链接"
                         onClick={() => {
                           navigator.clipboard.writeText(ReqStore.shareLink)
                           notifySuccess('分享链接', '复制成功');
                         }}
          />
          <OperAnswerBtn icon={<RefreshIcon/>}
                         hoverContent="重新生成"
                         onClick={() => {
                           ReqStore.queryQuestion(
                             BaseStore.question,
                             BaseStore.lang.query,
                             BaseStore.field.field,
                             BaseStore.fieldSpec.query);
                         }}
          />
        </div>

      }
    </>
  )
})

function regularizeMarkdown(markdown: string): string {
  // Pattern to match code blocks
  const codeBlockPattern = /```[\s\S]*?```/g;
  let lastIndex = 0;
  let result = '';

  // Function to replace \n with \n\n
  const replaceNewLines = (text: string) => text.replace(/\n(?!\n)/g, '\n\n');

  // Iterate over code blocks
  markdown.replace(codeBlockPattern, (match, index) => {
    // Process text outside code blocks
    result += replaceNewLines(markdown.slice(lastIndex, index));
    // Append code block without changes
    result += match;
    lastIndex = index + match.length;
    return match;
  });

  // Process any remaining text after the last code block
  result += replaceNewLines(markdown.slice(lastIndex));

  return result;
}