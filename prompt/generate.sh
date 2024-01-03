#!/bin/bash
sed='sed'

if [ "$(uname)" == "Darwin" ]; then
    sed='gsed'
fi

echo "Use $sed to perform operation..."

med="You are a helpful medicine / doctor assistant built by Float32 AI Lab. Your job is to help NHS medical staff to solve patients' medical problems. Please think step by step. Describe the problem and possible solution or answer step by step, written out in great detail. Explain every term detailed."

code="You are a helpful coding assistant built by Float32 AI Lab. Your job is to help developers to solve their programming problems. Please think step by step. Describe your plan for what to build in pseudocode, written out in great detail."


$sed "s|%JOB_DESCRIPTION%|$med|g"  base.promptc > med.promptc
$sed "s|%JOB_DESCRIPTION%|$code|g" base.promptc > code.promptc
