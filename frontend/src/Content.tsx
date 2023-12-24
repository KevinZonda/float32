import {Skeleton} from "tdesign-react";
import MarkdownPreview from '@uiw/react-markdown-preview';

export function Content() {
  const isLoading = false;
  if (isLoading) {
    return (
      <>
        <Skeleton animation={'flashed'} theme={'paragraph'} style={{paddingTop: '16px', paddingBottom: '16px'}}>
          <p>LOAD</p>
        </Skeleton>
      </>
    )
  }
  const md = source.split('\n').join('\n\n')
  // TODO LaTeX
  return (
    <>
      <h3 style={{paddingBottom: '16px', marginBlock: 0, marginBlockStart: 0, marginBlockEnd: 0, textAlign: 'left'}}>🔍 Answer</h3>
      <MarkdownPreview style={{textAlign: 'left', fontFamily: 'Linux Libertine'}} source={md} />
    </>
  )

}
const source = `React和Vue都是用于构建用户界面的流行JavaScript库/框架。虽然React通常被称为库，而Vue被称为框架，但它们都提供了类似的渐进式解决方案来构建复杂的应用程序。
React和Vue之间的一个关键区别是它们的渲染方式。React使用虚拟DOM（文档对象模型），它是实际DOM的轻量级副本。当数据发生变化时，React会更新虚拟DOM，并高效地确定更新实际DOM所需的最小更改数量。这种方法可以实现高效的渲染和更好的性能。
另一方面，Vue使用响应式数据模型。它利用了一种称为双向数据绑定的技术，其中数据的变化会自动反映在视图中，反之亦然。Vue通过使用一个响应式系统来跟踪依赖关系并相应地更新视图。这种方法简化了开发过程，使应用程序的状态更易于管理。
React和Vue之间的另一个区别是它们对数据的处理方式。React推崇使用不可变数据，即数据不直接修改，而是用新版本替换。这使得React可以通过比较新旧数据版本来高效地确定是否需要重新渲染组件。相比之下，Vue使用响应式数据，即对数据的更改会自动检测并触发受影响组件的重新渲染。这样可以实现更精确的更新，但可能会带来一些性能开销。
在语法方面，React使用JSX（JavaScript XML），允许开发人员直接在JavaScript中编写类似HTML的代码。JSX是一个强大的功能，可以在标记中使用JavaScript表达式。另一方面，Vue使用模板，它们基于HTML，并提供了一种声明性的方式来定义组件的结构和行为。
React和Vue都拥有丰富的生态系统，包括广泛的文档、活跃的社区以及各种第三方库和工具。它们都被广泛采用，并已被证明在构建现代Web应用程序方面非常有效。
总之，React和Vue都是构建用户界面的强大工具。React的虚拟DOM和对不可变数据的关注使其高效且性能优越，而Vue的响应式数据模型和基于模板的语法提供了简单性和易用性。选择React还是Vue最终取决于项目的具体要求和开发团队的偏好。
`;
