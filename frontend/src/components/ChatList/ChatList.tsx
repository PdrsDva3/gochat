import React, { useState, useEffect } from 'react';
import { List, Input, Button, Modal} from 'antd';
import { PlusOutlined, ReloadOutlined } from '@ant-design/icons';

interface Chat {
  id: number;
  name: string;
  avatarUrl?: string;
  lastMessage?: string;
}

const ChatList = () => {
  const [allChats, setAllChats] = useState<Chat[]>([
    { id: 1, name: 'Alice', avatarUrl: 'https://example.com/avatar1.jpg', lastMessage: 'Привет!' },
    { id: 2, name: 'Bob', avatarUrl: 'https://example.com/avatar2.jpg', lastMessage: 'Как дела?' },
    { id: 3, name: 'Charlie', avatarUrl: 'https://example.com/avatar3.jpg', lastMessage: 'Отлично!' }
  ]);
  const [filteredChats, setFilteredChats] = useState<Chat[]>([]);
  const [newChatName, setNewChatName] = useState('');
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    setFilteredChats(allChats);
  }, [allChats]);

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

  // Функция для добавления нового чата
  const addChat = () => {
    if (newChatName.trim()) {
      setAllChats(prevChats => [...prevChats, { id: prevChats.length + 1, name: newChatName, avatarUrl: `https://example.com/avatar${prevChats.length + 1}.jpg`, lastMessage: 'Новое сообщение' }]);
      setNewChatName('');
      setIsModalOpen(false);
    }
  };

  return (
    <>
    <Input
      placeholder="Поиск по имени чата"
      value={searchTerm}
      onChange={handleSearchChange}
      style={{ width: '200px', marginBottom: '16px' }}
    />
    <Button type="primary" onClick={handleClearSearch} icon={<ReloadOutlined />} style={{ marginLeft: '8px' }}>
      Очистить
    </Button>

    <Modal
      open={isModalOpen}
      footer={null}
      onCancel={() => setIsModalOpen(false)}
    >
      <Input
        placeholder="Введите имя чата"
        value={newChatName}
        onChange={(e) => setNewChatName(e.target.value)}
      />
      <Button type="primary" onClick={addChat} style={{ marginTop: '8px' }}>
        Добавить
      </Button>
    </Modal>

    <Button
      icon={<PlusOutlined />}
      onClick={() => setIsModalOpen(true)}
      style={{ marginBottom: '16px' }}
    >
      Добавить чат
    </Button>

    <div style={{ position: 'relative', marginTop: '16px' }}>
      <List
        itemLayout="horizontal"
        dataSource={filteredChats}
        renderItem={(item) => (
          <List.Item>
            <List.Item.Meta
              avatar={<img width={64} height={64} src={item.avatarUrl || '/defaultAvatar.png'} alt={item.name} />}
              title={<span style={{ fontWeight: 'bold' }}>{item.name}</span>}
              description={item.lastMessage}
            />
          </List.Item>
        )}
      />
    </div>
  </>
);
};

export default ChatList;
