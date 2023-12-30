import './App.css'
import {Button, Dropdown, Input} from "tdesign-react";
import {CodeIcon, Dividers1Icon, EarthIcon, InfoCircleIcon} from 'tdesign-icons-react';
import React from "react";
import ReqStore from "./Store/ReqStore.ts";
import {observer} from "mobx-react-lite";
import {useNavigate, useParams} from "react-router-dom";
import {ContentLayout} from "./ContentLayout.tsx";
import {FcPlus} from "react-icons/fc";
import {MdOutlineMedicalInformation} from "react-icons/md";

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
  field: string
  options: { content: string, value: { query: string, content: string } }[]
  icon: React.ReactElement
  subIcon: React.ReactElement
}

const fieldsOpt = [
  {
    content: '程序开发',
    value: {
      content: '程序开发',
      field: 'code',
      options: progLangOpt,
      icon: <Dividers1Icon size={"16"}/>,
      subIcon: <CodeIcon size="16"/>
    },
  },
  {
    content: '医学',
    value: {
      content: '医学',
      field: 'med',
      options: [
        newQuery('NHS (UK)', 'nhs'),
        newQuery('NICE (UK)', 'nice'),
        newQuery('CDC (US)', 'cdc'),
        newQuery('默认', ''),
      ],
      icon: <div style={{paddingRight: '8px'}}>
        <FcPlus size={'18px'} style={{
          fontSize: '17px',
          verticalAlign: 'middle',
          marginBottom: '2px'
        }} class={'t-icon'}/></div>,
      subIcon: <div style={{paddingRight: '8px'}}>
        <MdOutlineMedicalInformation size={'18px'} style={{
          fontSize: '17px',
          verticalAlign: 'middle',
          marginBottom: '2px'
        }} class={'t-icon'}/></div>
    },
  }
]

const dropdownBtnStyle = {paddingRight: '8px', paddingLeft: '8px'}

export const App = observer(() => {
  const [lang, setLang] = React.useState(langOpt[0].value);
  const [field, setField] = React.useState(fieldsOpt[0].value);
  const [fieldSpec, setFieldSpec] = React.useState(field.options[0].value);
  const [fieldIcon, setFieldIcon] = React.useState(field.icon);
  const [subIcon, setSubIcon] = React.useState(field.subIcon);

  const query = useParams();
  const id = query.id
  if (id && id !== '') {
    ReqStore.queryHistory(id);
  }

  const nav = useNavigate()
  return (
    <>
      <div style={{width: '100%', textAlign: 'center'}} onClick={() => {
        ReqStore.isRainbow = !ReqStore.isRainbow
      }}>
        <h1 style={{
          fontFamily: `'PT Sans Narrow', sans-serif;`,
          color: 'black',
          marginTop: ReqStore.currentAns && 0,
          width: 'fit-content',
          marginLeft: 'auto',
          marginRight: 'auto',
          paddingLeft: '8px',
          paddingRight: '8px',
          marginBottom: '12px',
          borderRadius: '8px',
        }}
            className={ReqStore.isRainbow ? 'rainbow' : ''}
        >
          <span style={{fontFamily: `'PT Sans', sans-serif`, fontStyle: 'italic'}}>float32 AI</span>
          <span style={{fontFamily: `'PT Sans Narrow', sans-serif`, fontStyle: 'italic'}}>: Search Done Right</span>
        </h1>
      </div>


      <Input
        placeholder="请输入你的问题"
        size="large"
        value={ReqStore.question}
        onChange={(e) => {
          ReqStore.question = e
        }}
        onEnter={(question, e) => {
          if (!e.e.nativeEvent.isComposing && question !== '') {
            ReqStore.queryQuestion(question, lang.query, field.field, fieldSpec.query);
          }
        }}
      />
      <Dropdown options={langOpt} onClick={(e) => setLang(e.value as { query: string, content: string })}>
        <Button style={dropdownBtnStyle} variant="text" icon={<EarthIcon size="16"/>}>
          {lang.content}
        </Button>
      </Dropdown>
      {
        fieldsOpt.length > 1 &&
          <Dropdown options={fieldsOpt} onClick={(e) => {
            const f = e.value as IField
            setFieldIcon(f.icon)
            setSubIcon(f.subIcon)
            setField(f)
            setFieldSpec(f.options[0].value)
            if (f.field === 'med') {
              ReqStore.warning = '本网站上的信息来自互联网或人工智能生成内容。网站无意提供医疗建议，也不能代替由资质的医生、药剂师或其他医疗保健专业人员所提供的咨询。读者不能因为本网站上提供的某些信息，而无视医生的建议或延迟就医。'
            } else {
              ReqStore.warning = ''
            }
          }}>
              <Button style={dropdownBtnStyle} variant="text" icon={fieldIcon}>
                {field.content}
              </Button>
          </Dropdown>
      }
      {
        field.options.length > 1 &&
          <Dropdown options={field.options}
                    onClick={(e) => setFieldSpec(e.value as { query: string, content: string, icon: string })}>
              <Button style={dropdownBtnStyle} variant="text" icon={subIcon}>
                {fieldSpec.content}
              </Button>
          </Dropdown>
      }
      <Button style={dropdownBtnStyle} theme="default" variant="text" icon={<InfoCircleIcon size="16"/>} onClick={() => {
        nav('/about')
      }}>
        关于
      </Button>
      <div style={{height: '16px'}}></div>
      <ContentLayout/>
    </>
  )
})

