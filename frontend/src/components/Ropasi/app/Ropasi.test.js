import React from 'react';
import ReactDOM from 'react-dom';
import Ropasi from './Ropasi';

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<Ropasi />, div);
  ReactDOM.unmountComponentAtNode(div);
});
