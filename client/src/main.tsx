import { hydrateRoot } from 'react-dom/client';
import App from './App.tsx';
import './index.css';
import { BrowserRouter } from 'react-router-dom';

hydrateRoot(document.getElementById('root')!, (
  <BrowserRouter>
    <App />
  </BrowserRouter>
));
