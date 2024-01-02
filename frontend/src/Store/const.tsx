import React from "react";
import {FcPlus} from "react-icons/fc";
import { MdOutlineMedicalInformation} from "react-icons/md";
import {RiCodeSSlashLine, RiCompasses2Line} from "react-icons/ri";

function newQuery(content: string, query: string = content) {
  return {
    content: content,
    value: {
      query: query,
      content: content
    },
  }
}

export const langOpt = [
  newQuery('简体中文', 'zh'),
  newQuery('English', 'en'),
];


export const progLangOpt = [
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

export interface IField {
  content: string
  field: string
  options: { content: string, value: { query: string, content: string } }[]
  icon: React.ReactElement
  subIcon: React.ReactElement
}

export const fieldsOpt = [
  {
    content: '程序开发',
    value: {
      content: '程序开发',
      field: 'code',
      options: progLangOpt,
      icon: <div style={{paddingRight: '5px'}}>
        <RiCompasses2Line size={'18px'} style={{
          fontSize: '17px',
          verticalAlign: 'middle',
          marginBottom: '2px',
        }} class={'t-icon'}/></div>,
      subIcon: <div style={{paddingRight: '5px'}}>
        <RiCodeSSlashLine size={'18px'} style={{
          fontSize: '17px',
          verticalAlign: 'middle',
          marginBottom: '2px',
        }} class={'t-icon'}/></div>
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
      icon: <div style={{paddingRight: '5px'}}>
        <FcPlus size={'18px'} style={{
          fontSize: '17px',
          verticalAlign: 'middle',
          marginBottom: '2px'
        }} class={'t-icon'}/></div>,
      subIcon: <div style={{paddingRight: '5px'}}>
        <MdOutlineMedicalInformation size={'18px'} style={{
          fontSize: '17px',
          verticalAlign: 'middle',
          marginBottom: '2px'
        }} class={'t-icon'}/></div>
    },
  }
]
