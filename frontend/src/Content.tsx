import {Skeleton} from "tdesign-react";
import MarkdownPreview from '@uiw/react-markdown-preview';
import ReqStore from "./Store/ReqStore.ts";
import {observer} from "mobx-react-lite";
import './Warning.css'

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
            <p style={{textAlign: 'left', paddingBottom: 0, marginBlockEnd:0, marginBottom:0}}>
              {ReqStore.warning}
            </p>
          </div>


        </div>
      </>
  )
})

export const Content = observer(() => {
  // TODO LaTeX
  if (!ReqStore.isLoading && ReqStore.currentAns === '') {
    return <>
      <Warning/>
    </>
  }
  return (
    <>
    <Warning/>
      <h3 style={{
        paddingBottom: '16px',
        marginBlock: 0,
        marginBlockStart: 0,
        marginBlockEnd: 0,
        textAlign: 'left'
      }}>{ReqStore.isFailed ? '⚠️ Error' : '🔍 Answer'}</h3>
      <MarkdownPreview
        wrapperElement={{
          "data-color-mode": "light"
        }}
        style={{textAlign: 'left', fontFamily: 'Linux Libertine'}}
        source={regularizeMarkdown(ReqStore.currentAns)}/>
      {ReqStore.isLoading ?
        <Skeleton animation={'flashed'} theme={'paragraph'} style={{paddingTop: '16px', paddingBottom: '16px'}}>
          <p>LOAD</p>
        </Skeleton> : null}
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