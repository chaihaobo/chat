import { FC, useState } from "react";
import { Input, Button, Popover } from "antd";
import { SendOutlined, SmileOutlined } from "@ant-design/icons";
import styled from "styled-components";
import data from '@emoji-mart/data';
import Picker from '@emoji-mart/react';

interface ChatInputProps {
    className?: string;
    onSendMessage: (message: string) => void;
}

const ChatInput: FC<ChatInputProps> = ({ className, onSendMessage }) => {
    const [message, setMessage] = useState("");
    
    const handleEmojiSelect = (emoji: any) => {
        setMessage(prev => prev + emoji.native);
    };

    const handleSend = () => {
        if (message.trim()) {
            onSendMessage(message);
            setMessage("");
        }
    };

    const handleKeyPress = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
        if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            handleSend();
        }
    };

    return (
        <div className={className}>
            <Input.TextArea
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                onKeyPress={handleKeyPress}
                placeholder="输入消息..."
                autoSize={{ minRows: 2, maxRows: 4 }}
            />
            <Popover
                content={
                    <Picker
                        data={data}
                        onEmojiSelect={handleEmojiSelect}
                        locale="zh"
                        theme="light"
                        previewPosition="none"
                    />
                }
                trigger="click"
                placement="topRight"
            >
                <Button
                    icon={<SmileOutlined />}
                    style={{ marginRight: 8 }}
                />
            </Popover>
            <Button 
                type="primary"
                icon={<SendOutlined />}
                onClick={handleSend}
                disabled={!message.trim()}
            >
                发送
            </Button>
        </div>
    );
};

export default styled(ChatInput)`
    display: flex;
    gap: 10px;
    padding: 20px;
    background-color: white;
    border-top: 1px solid #e8e8e8;
    width: 70%;

    .ant-input-textarea {
        flex: 1;
    }

    .ant-btn {
        align-self: flex-end;
    }
`;