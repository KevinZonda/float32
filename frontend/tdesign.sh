#!/bin/bash
cd public
rm tdesign.min.css
rm tdesign.min.js
wget https://unpkg.com/tdesign-react/dist/tdesign.min.css
wget https://unpkg.com/tdesign-react/dist/tdesign.min.js
cd ..