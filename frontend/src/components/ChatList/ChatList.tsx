import React, { useState, useEffect } from 'react';
import { List, Input, Button, Modal } from 'antd';
import { PlusOutlined, ReloadOutlined } from '@ant-design/icons';
import './ChatList.scss';
import { saveChats, loadChats } from './utils.ts';
interface Chat {
  id: number;
  name: string;
  avatarUrl?: string;
  lastMessage?: string;
}

const ChatList: React.FC = () => {
  const [allChats, setAllChats] = useState<Chat[]>([
		{
			id: 1,
			name: 'Alice',
			avatarUrl: 'https://example.com/avatar1.jpg',
			lastMessage: 'Привет!',
		},
		{
			id: 2,
			name: 'Bob',
			avatarUrl: 'https://example.com/avatar2.jpg',
			lastMessage: 'Как дела?',
		},
		{
			id: 3,
			name: 'Charlie',
			avatarUrl: 'https://example.com/avatar3.jpg',
			lastMessage: 'Отлично!',
		},
	]);
	const [filteredChats, setFilteredChats] = useState<Chat[]>([]);
  const [newChatName, setNewChatName] = useState('');
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    setFilteredChats(allChats);
  }, [allChats]);

	useEffect(() => {
    // Загружаем чаты при монтировании компонента
    const loadedChats = loadChats();
    setAllChats(loadedChats);
  }, []);

  const filterChats = (chats: Chat[], searchTerm: string): Chat[] => {
    return chats.filter((chat: Chat) =>
      chat.name.toLowerCase().includes(searchTerm.toLowerCase())
    );
  };

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const filteredChats = filterChats(allChats, e.target.value);
    setFilteredChats(filteredChats);
    setSearchTerm(e.target.value);
  };

  const handleClearSearch = () => {
    setSearchTerm('');
    setFilteredChats(allChats);
  };

  const addChat = () => {
    if (newChatName.trim()) {
      setAllChats((prevChats) => [
        ...prevChats,
        {
          id: prevChats.length + 1,
          name: newChatName,
          avatarUrl: `https://example.com/avatar${prevChats.length + 1}.jpg`,
          lastMessage: 'Новое сообщение',
        },
      ]);
			saveChats(allChats);
      setNewChatName('');
      setIsModalOpen(false);
    }
  };

  const handleChatClick = (chatId: number) => {
    window.location.href = `/chat/${chatId}`;
    console.log(`Открыть чат с ID: ${chatId}`);
		window.history.pushState({}, '', `/chat/${chatId}`);
  };

	return (
		<>
			<div className="chatList">
				<div className="some-container">
					<Input
						placeholder="Поиск по имени чата"
						value={searchTerm}
						onChange={handleSearchChange}
						style={{ width: '100%', marginBottom: '16px' ,}}
					/>
					<Button
						style={{width:"200px", marginBottom:"16px"}}
						type="primary"
						onClick={handleClearSearch}
						icon={<ReloadOutlined />}
					>
						Очистить
					</Button>

					<Button
						icon={<PlusOutlined />}
						onClick={() => setIsModalOpen(true)}
						style={{width:"200px", marginBottom:"16px"}}
					>
						Добавить чат
					</Button>
				</div>

				<Modal
					style={{}}
					open={isModalOpen}
					footer={null}
					onCancel={() => setIsModalOpen(false)}
				>
					<Input
						style={{ marginTop: '27px',}}
						placeholder="Введите имя чата"
						value={newChatName}
						onChange={(e) => setNewChatName(e.target.value)}
					/>
					<Button type="primary" onClick={addChat} style={{ marginTop: '16px' }}>
						Добавить
					</Button>
				</Modal>

				<div style={{ position:'relative', marginTop: '16px'}}>
          <List
            itemLayout="horizontal"
            dataSource={filteredChats}
            renderItem={(item) => (
              <List.Item onClick={() => handleChatClick(item.id)}>
                <List.Item.Meta
                  avatar={
                    <img
                      width={64}
                      height={64}
                      src={item.avatarUrl || '/defaultAvatar.png'}
                      alt={item.name}
                    />
                  }
                  title={<span style={{ fontWeight: 'bold' }}>{item.name}</span>}
                  description={item.lastMessage}
                />
              </List.Item>
            )}
          />
        </div>
      </div>
    </>
  );
};

export default ChatList;
