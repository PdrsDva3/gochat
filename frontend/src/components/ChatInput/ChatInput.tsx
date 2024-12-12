// нижняя панель ввода приложений
import { Button, Input } from 'antd';
import './Chatinput.scss';
import {PaperClipOutlined, SmileOutlined, AudioOutlined} from '@ant-design/icons';




const ChatInput = () => {

	const handleFileClick = () => {
    console.log('Клик по кнопке загрузки файла');
    // Здесь можно добавить логику для открытия диалога выбора файла
  };

  const handleEmojiClick = () => {
    console.log('Клик по кнопке эмодзи');
    // Здесь можно добавить логику для открытия панели с эмодзи
  };

	const handleVoiceNoteClick = () => {
    console.log('Начало записи голосового сообщения');
    // Здесь можно добавить логику для начала записи
  };



	return (
    <div className="input-container">
      {/* Кнопка для загрузки файла */}
      <Button icon={<PaperClipOutlined/>} onClick={handleFileClick}/>

      {/* Поле ввода сообщения */}
      <Input
        placeholder="Введите сообщение..."
        style={{ flex: 1 }}
        onPressEnter={(e) => {
          e.preventDefault();
          console.log('Отправка сообщения');
          // Здесь можно добавить логику отправки сообщения
        }}
      />

      {/* Кнопка для добавления эмодзи */}
      <Button icon ={<SmileOutlined />} onClick={handleEmojiClick}/>

			{/* Кнопка для записи голосовых сообщений */}
      <Button icon={<AudioOutlined />} onClick={handleVoiceNoteClick}/>
    </div>
  );
};

export default ChatInput;
