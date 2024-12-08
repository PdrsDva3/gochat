import React, { useState, useEffect} from 'react';
import { Layout } from 'antd';
import { Input, Button, Avatar, List } from 'antd';
import { UserOutlined, EllipsisOutlined, SettingOutlined } from '@ant-design/icons';
import { message } from 'antd';
import { saveMessages, loadMessages } from '../ChatList/utils.ts';

const { Header, Content, Footer } = Layout;

interface Message {
	id: number;
	sender: string;
	content: string;
	timestamp: Date;
}

interface User {
	name: string;
	avatar: string;
}

interface ChatWindowProps {
	chatId: number;
	initialMessages?: Message[];
	initialUser?: User;
}

const ChatWindow: React.FC<ChatWindowProps> = ({ chatId, initialMessages, initialUser }) => {
	const [messages, setMessages] = useState<Message[]>(initialMessages || []);
	const [inputValue, setInputValue] = useState('');
	const [currentUser, setCurrentUser] = useState<User>(
		initialUser || { name: '', avatar: '' },
	);

	useEffect(() => {
    // Загружаем сообщения для выбранного чата при монтировании компонента
    const loadedMessage = loadMessages(chatId);
    if (loadedMessage) {
      setMessages([loadedMessage]);
    }
  }, [chatId]);

	const sendMessage = async () => {
		if (!inputValue.trim()) return;

		const newMessage: Message = {
			id: Date.now(),
			sender: currentUser.name,
			content: inputValue,
			timestamp: new Date(),
		};

		setMessages((prevMessages) => [...prevMessages, newMessage]);

		// Отображаем уведомление после отправки сообщения
		message.success('Сообщение успешно отправлено!', 1.5);

		console.log('Отправлено сообщение:', newMessage.content);

		saveMessages(newMessage);
		setInputValue('');
	};

	const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		setInputValue(e.target.value);
	};

	const updateUser = (newUserData: Partial<User>) => {
		setCurrentUser((prevUser) => ({
			...prevUser,
			...newUserData,
		}));
	};

	return (
		<Layout style={{ height: '100vh' }}>
			<Header style={{ display: 'flex', padding: '16px', backgroundColor: '#1890ff', color: 'white', height:"64px" }}>
				<div
					style={{
						display: 'flex',
						alignItems: 'center',
						justifyContent: 'space-between',
						marginRight: '16px',
					}}
				>
					<div style={{ display: 'flex', alignItems: 'center' }}>
						{currentUser.avatar ? (
							<Avatar size={32} src={currentUser.avatar} alt={currentUser.name} />
						) : (
							<Avatar size={32} icon={<UserOutlined />} />
						)}
						<span style={{ marginLeft: '8px' }}>{currentUser.name}</span>
					</div>
					<Button
						icon={<SettingOutlined />}
						onClick={() =>
							updateUser({
								name: 'Новый пользователь',
								avatar: 'https://example.com/new-avatar.jpg',
							})
						}
					/>
				</div>
			</Header>

			<Content
				style={{
					padding: '24px',
					overflowY: 'auto',
					maxHeight: 'calc(100vh-160px)',
					border: '1px solid #e8e8e8',
				}}
			>
				<List
					itemLayout="vertical"
					size="large"
					dataSource={messages}
					renderItem={(item) => (
						<List.Item>
							<List.Item.Meta
								avatar={
									item.sender === 'Bot' ? (
										<Avatar icon="robot" />
									) : (
										<Avatar
											src={
												item.sender === currentUser.name ? currentUser.avatar : undefined
											}
										/>
									)
								}
								description={item.content}
								title={`${item.sender}: ${item.timestamp.toLocaleString()}`}
							/>
						</List.Item>
					)}
				/>
			</Content>
			<Footer style={{ padding: '16px', textAlign: 'right' }}>
				<Input
					value={inputValue}
					onChange={handleInputChange}
					onPressEnter={sendMessage}
					style={{ width: '70%', marginRight: '8px'}}
				/>
				<Button type="primary" onClick={sendMessage}>
					Отправить
				</Button>
				<Button icon={<EllipsisOutlined />} />
			</Footer>
		</Layout>
	);
};

export default ChatWindow;
