import './App.css'
import {Button, Dropdown, Input} from "tdesign-react";
import React, {useEffect, useState} from "react";
import ReqStore from "./Store/ReqStore.ts";
import {observer} from "mobx-react-lite";
import {ContentLayout} from "./ContentLayout.tsx";
import {langOpt, fieldsOpt, IField, ReactIcon} from "./Store/const.tsx";
import {BaseStore} from "./Store/BaseStore.ts";
import {RiQuestionAnswerLine} from "react-icons/ri";
import {Conditional, useQuery} from "./CommonComponents.tsx";
import {FaGithub} from "react-icons/fa";

const dropdownBtnStyle = {paddingRight: '8px', paddingLeft: '8px'}

export const App = observer(() => {
  const [isMobile, setIsMobile] = useState(window.innerWidth < 720)
  useEffect(() => {
    window.addEventListener("resize", () => setIsMobile(window.innerWidth < 720))
  })

  const [lang, setLang] = React.useState(langOpt[0].value);
  const [field, setField] = React.useState(fieldsOpt[0].value);
  const [fieldSpec, setFieldSpec] = React.useState(field.options[0].value);
  const [fieldIcon, setFieldIcon] = React.useState(field.icon);
  const [subIcon, setSubIcon] = React.useState(field.subIcon);

  const query = useQuery();
  const id = query.get('id')
  if (id && id !== '' && id !== ReqStore.shareId) {
    ReqStore.queryHistory(id);
  } else {
    const q = (query.get('q') ?? '').trim()
    if (q && q !== '' && q !== ReqStore.question) {
      const field = query.get('field') ?? 'code' // field
      const lang = query.get('lang') ?? 'zh' // lang
      const spec = query.get('spec') ?? '' // spec
      BaseStore.fieldSpec = fieldSpec
      BaseStore.question = q
      ReqStore.question = q
      // TODO: sync field & spec to state
      ReqStore.queryQuestion(q, lang, field, spec);
    }
  }

  return (
    <>
      <div style={{width: '100%', textAlign: 'center'}}>
        <h1 style={{
          fontFamily: `'PT Sans Narrow', sans-serif`,
          color: 'black',
          marginTop: ReqStore.currentAns && 0,
          width: 'fit-content',
          marginLeft: 'auto',
          marginRight: 'auto',
          paddingLeft: '8px',
          paddingRight: '8px',
          marginBottom: '12px',
          borderRadius: '8px',
          cursor: 'pointer',
          userSelect: 'none'
        }}
            className={ReqStore.isRainbow ? 'rainbow' : ''}
            onClick={() => {
              console.log(window.location.pathname)
              if (window.location.pathname !== '/') {
                window.location.href = '/'
              } else {
                ReqStore.isRainbow = !ReqStore.isRainbow
              }
            }}
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
          if (e.e.nativeEvent.isComposing || question === '') {
            return
          }
          BaseStore.lang = lang
          BaseStore.field = field
          BaseStore.fieldSpec = fieldSpec
          BaseStore.question = question
          ReqStore.queryQuestion(question, lang.query, field.field, fieldSpec.query);

        }}
      />
      <Dropdown options={langOpt} onClick={(e) => {
        const l = e.value as {
          query: string, content: string
        }
        setLang(l)
      }}>
        <Button style={dropdownBtnStyle} variant="text"
                icon={<ReactIcon><RiQuestionAnswerLine/></ReactIcon>}>
          {lang.content}
        </Button>
      </Dropdown>

      <Conditional condition={fieldsOpt.length > 1}>
        <Dropdown options={fieldsOpt} onClick={(e) => {
          const f = e.value as IField
          setFieldIcon(f.icon)
          setSubIcon(f.subIcon)
          setField(f)
          setFieldSpec(f.options[0].value)
          ReqStore.warning =
            f.field === 'med'
              ? '本网站上的信息来自互联网或人工智能生成内容。网站无意提供医疗建议，也不能代替由资质的医生、药剂师或其他医疗保健专业人员所提供的咨询。读者不能因为本网站上提供的某些信息，而无视医生的建议或延迟就医。'
              : ''
        }}>
          <Button style={dropdownBtnStyle} variant="text" icon={fieldIcon}>
            {field.content}
          </Button>
        </Dropdown>
      </Conditional>

      <Conditional condition={field.options.length > 1}>
        <Dropdown options={field.options}
                  onClick={(e) => {
                    const fs = e.value as { query: string, content: string, icon: string }
                    setFieldSpec(fs)
                  }}>
          <Button style={dropdownBtnStyle} variant="text" icon={subIcon}>
            {fieldSpec.content}
          </Button>
        </Dropdown>
      </Conditional>

      <Conditional condition={!isMobile}>
        <Button style={dropdownBtnStyle} theme="default" variant="text"
                icon={<ReactIcon><FaGithub/></ReactIcon>}
                href="https://github.com/KevinZonda/float32">
          GitHub
        </Button>
      </Conditional>

      <div style={{height: '16px'}}></div>
      <ContentLayout/>
    </>
  )
})

