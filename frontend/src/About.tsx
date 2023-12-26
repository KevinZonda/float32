import './App.css'
import {useNavigate} from "react-router-dom";

export const About = () => {
  const nav = useNavigate()
  return (
    <>
      <h1 onClick={() => nav('/')}>About :: Float32.app
      </h1>

      <div style={{textAlign: 'left'}}>
        <p>
          Float32.App 是一个文档搜索引擎。其的设计目的是为了帮助开发人员优雅且快速找到可靠的文档。
        </p>
        <h2>可靠性</h2>
        Float32.App 通过引入外部知识以对抗大语言模型幻觉问题。
        当你每次请求时，Float32 Agent 都会通过网页信息用于提升结果可靠性。
        相比于传统 ChatGPT prompt Engineering来说，这种操作会花费更多时间，但是会带来不少好处，包括但不局限于
        <ul>
          <li>ChatGPT有了可信来源，因此不会再胡说八道了</li>
          <li>ChatGPT无需草稿纸（ToC技术）就可以实现更好的效果。因此回答也更简短</li>
        </ul>
        <h2>时效性</h2>
        <p>
          每次请求 Float32.App 时，我们的 Agent 都会即时抓取网页信息用于提升时效性。
        </p>
        <h2>结语</h2>
        <p>
          3天，我们从 0 构建了整个 Float32.App。
          <br></br>
          无意间回首，开发工具的复杂度和易用性都得到了质的飞跃。
          但是开发人员的心智负担却越来越重了。大家需要了解更多的语言，框架，工具，设计复杂的系统。
        </p>
        <p>愿大家永远开心。Merry Christmas! 🎄</p>
        <p>Docs Done Right.<br/>
          Float32.App Team in UK<br/>
          25/Dec/2023
        </p>
      </div>
    </>
  )
}