import './App.css'
import {Button, Dropdown, Input} from "tdesign-react";
import { Icon } from 'tdesign-icons-react';
import React from "react";
import {Content} from "./Content.tsx";

const langOpt = [
  {
    content: 'English',
    value: 'en',
  },
  {
    content: '简体中文',
    value: 'zh',
  }
];

function langMap(lang: string) {
  switch (lang) {
    case 'en':
      return 'English';
    case 'zh':
      return '简体中文';
    default:
      return 'Default';
  }
}


function App() {
  const [lang, setLang] = React.useState('zh');
  return (
    <>
      <h1 style={{fontFamily: 'Linux Libertine'}}><span style={{fontStyle: 'italic'}}>float32</span> AI : Docs 	&times; Elegant</h1>
      <p className="read-the-docs">
        曲径通幽，拒绝繁琐文档，告别无效思考。清雅绝俗，挑战世俗之见。
      </p>
      <Input
        placeholder="请输入你的问题"
        size="large"
      />
      <Dropdown options={langOpt} onClick={(e)=> setLang(e.value as string)}>
        <Button variant="text" suffix={<Icon name="chevron-down" size="16" />}>
          {langMap(lang)}
        </Button>
      </Dropdown>
      <div style={{height: '16px'}}></div>
      <Content />
    </>
  )
}

export default App
