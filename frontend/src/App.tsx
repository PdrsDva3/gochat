import { BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import { Row, Col } from 'antd';
import ChatList from './components/ChatList/ChatList';
import ChatWindow from './components/ChatWindow/ChatWindow';
import './styles/App.scss';

export const App = () => {
  return (
    <Router>
      <Row>
        <Col span={18} push={6}>
          <Routes>
            <Route path="/" element={<ChatList />} />
            <Route path="/chat/:id" element={<ChatWindow />} />
          </Routes>
        </Col>
        <Col span={6} pull={18}>
          <ChatList />
        </Col>
      </Row>
    </Router>
  );
};
