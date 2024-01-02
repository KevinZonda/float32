import {useLocation} from "react-router-dom";
import React from "react";

export function useQuery() {
  const {search} = useLocation();

  return React.useMemo(() => new URLSearchParams(search), [search]);
}

export const Conditional = ({condition, children}: { condition: boolean, children: React.ReactElement }) => {
  if (condition) {
    return children
  }
  return <></>
}
