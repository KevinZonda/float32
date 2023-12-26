import './App.css'
import {Button, Dropdown, Input} from "tdesign-react";
import {Icon} from 'tdesign-icons-react';
import React from "react";
import ReqStore from "./Store/ReqStore.ts";
import {observer} from "mobx-react-lite";
import {useNavigate} from "react-router-dom";
import {ContentLayout} from "./ContentLayout.tsx";

function newQuery(content: string, query: string = content) {
  return {
    content: content,
    value: {
      query: query,
      content: content
    },
  }
}

const langOpt = [
  newQuery('简体中文', 'zh'),
  newQuery('English', 'en'),
];


const progLangOpt = [
  newQuery('默认语言', ''),
  newQuery('Go', 'golang'),
  newQuery('Python', 'python'),
  newQuery('PyTorch', 'Python, Pytorch, Numpy'),
  newQuery('Rust', 'rust'),
  newQuery('JavaScript', 'JavaScript'),
  newQuery('TypeScript', 'TypeScript'),
  newQuery('Java', 'Java'),
  newQuery('C#', 'C#'),
  newQuery('C', 'C'),
  newQuery('C++', 'C++'),
  newQuery('Haskell', 'Haskell'),
];

interface IField {
  content: string
  options: { content: string, value: { query: string, content: string } }[]
  icon: React.ReactElement
}

const fieldsOpt = [
  {
    content: '程序开发',
    value: {
      content: '程序开发',
      options: progLangOpt,
      icon: <Icon name="code" size="16"/>
    },
  }
]


export const App = observer(() => {
  const [lang, setLang] = React.useState(langOpt[0].value);
  const [field, setField] = React.useState(fieldsOpt[0].value);
  const [progLang, setProgLang] = React.useState(field.options[0].value);
  const [fieldIcon, setFieldIcon] = React.useState(field.icon);
  const nav = useNavigate()
  return (
    <>
      <h1 style={
        ReqStore.currentAns === '' ? {fontFamily: `'Linux Libertine', 'Linux Libertine O'`} : {
          fontFamily: `'Linux Libertine', 'Linux Libertine O'`,
          marginTop: 0
        }
      }><span style={{fontStyle: 'italic'}}>float32</span> AI : Docs  &times; Elegant</h1>
      <p className="read-the-docs">
        曲径通幽，拒绝繁琐文档，告别无效思考。清雅绝俗，挑战世俗之见。
      </p>
      <Input
        placeholder="请输入你的问题"
        size="large"
        onEnter={(e) => {
          ReqStore.queryQuestion(e, lang.query, progLang.query);
        }}
      />
      <Dropdown options={langOpt} onClick={(e) => setLang(e.value as { query: string, content: string })}>
        <Button variant="text" icon={<Icon name="earth" size="16"/>}>
          {lang.content}
        </Button>
      </Dropdown>
      {
        fieldsOpt.length > 1 &&
          <Dropdown options={fieldsOpt} onClick={(e) => {
            const f = e.value as IField
            setFieldIcon(f.icon)
            setField(f)
          }}>
              <Button variant="text" icon={<Icon name="dividers-1" size="16"/>}>
                {field.content}
              </Button>
          </Dropdown>
      }
      {
        field.options.length > 1 &&
          <Dropdown options={field.options}
                    onClick={(e) => setProgLang(e.value as { query: string, content: string, icon: string })}>
              <Button variant="text" icon={fieldIcon}>
                {progLang.content}
              </Button>
          </Dropdown>
      }
      <Button theme="default" variant="text" icon={<Icon name="info-circle"/>} onClick={() => {
        nav('/about')
      }}>
        关于
      </Button>
      <div style={{height: '16px'}}></div>
      <ContentLayout/>
    </>
  )
})