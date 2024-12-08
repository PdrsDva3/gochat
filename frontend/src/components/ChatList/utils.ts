interface Chat {
  id: number;
  name: string;
  avatarUrl?: string;
  lastMessage?: string;
}

interface Message {
  id: number;
  sender: string;
  content: string;
  timestamp: Date;
}

const CHATS_KEY = 'chats';
const MESSAGES_KEY = 'messages';

export const saveChats = (chats: Chat[]): void => {
  localStorage.setItem(CHATS_KEY, JSON.stringify(chats));
};

export const loadChats = (): Chat[] => {
  const chatsData = localStorage.getItem(CHATS_KEY);
  return chatsData ? JSON.parse(chatsData) : [];
};

export const saveMessages = (message: Message): void => {
  const messages = JSON.parse(localStorage.getItem(MESSAGES_KEY) || '{}');
  messages[message.id] = message;
  localStorage.setItem(MESSAGES_KEY, JSON.stringify(messages));
};

export const loadMessages = (chatId: number): Message | null => {
  const messages = JSON.parse(localStorage.getItem(MESSAGES_KEY) || '{}');
  return messages[chatId];
};
