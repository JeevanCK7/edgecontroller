import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import 'typeface-roboto';
import App from './App';
require('dotenv').config()

ReactDOM.render(
  <App />
  , document.getElementById('root'));

if (module.hot) {
  module.hot.accept();
}