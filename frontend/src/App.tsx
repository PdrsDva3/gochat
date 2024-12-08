
import ChatList from './components/ChatList/ChatList';
import ChatWindow from './components/ChatWindow/ChatWindow';
import { Row, Col } from 'antd';
import './styles/App.scss';

export const App = () => {

	return (
		<>
			<Row gutter={[16, 16]}>
      <Col span={6}>
        <ChatList />
      </Col>
      <Col span={18}>
        <ChatWindow />
      </Col>
    </Row>
		</>
	);
};
