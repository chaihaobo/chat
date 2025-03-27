import {FC, useState} from "react";
import {styled} from "styled-components";
import Messagebox, {MessageBoxItem} from "./components/messagebox.tsx";
import ChatInput from "./components/chatinput.tsx";
import ContactList, {Contact} from "./components/contactlist.tsx";
import {message} from "antd";
import useChatConnection from "../../hooks/useChatConnection.ts";

enum EventType {
    SendMessage = 1,
    ReceiveMessage = 2,
}

interface Payload<T> {
    event: EventType
    data: T
}

interface ReceiveMessage {
    from: number
    content: string
}


const initialContacts: Contact[] = [];

interface ChatHistory {
    [contactId: number]: MessageBoxItem[];
}

const initialChatHistory: ChatHistory = {};

const Index: FC<{ className?: string }> = ({className}) => {
    const [messageApi, contextHolder] = message.useMessage();
    const [contacts, setContacts] = useState<Contact[]>(initialContacts);
    const [selectedContact, setSelectedContact] = useState<Contact | null>(contacts[0]);
    const [chatHistory, setChatHistory] = useState<ChatHistory>(initialChatHistory);

    const chatConnection = useChatConnection({
        onOpen: () => {
            messageApi.open({
                type: 'success',
                content: "连接成功",
            });
        },
        onFail: () => {
            messageApi.open({
                type: 'error',
                content: '连接失败，请重新输入凭证',
            });
        },
        onMessage: (e: MessageEvent<string>) => {
            console.log(e)
            messageApi.open({
                type: 'success',
                content: e.data,
            });
            onMessage(e.data)
        },
    })


    const handleReceiveMessage = (payload: Payload<ReceiveMessage>) => {
        const {from, content} = payload.data;

        // Create a new message
        const newMessage: MessageBoxItem = {
            senderID: from,
            senderAvatar: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png",
            senderName: "用户" + from,
            content,
        };

        // Update chat history
        setChatHistory(prev => ({
            ...prev,
            [from]: [...(prev[from] || []), newMessage]
        }));
        console.log(contacts)
        setContacts(prev => {
            let newContacts = [...prev]
            if (!newContacts.find(c => c.id === from)) {
                newContacts = [{
                    id: newMessage.senderID,
                    name: newMessage.senderName,
                    avatar: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png",
                    lastMessage: newMessage.content,
                    unreadCount: 1,
                }, ...newContacts]
                return newContacts
            }
            newContacts = newContacts.map(contact => {
                // 更新最新的消息
                if (contact.id === from) {
                    return {
                        ...contact,
                        lastMessage: contact.lastMessage,
                        unreadCount: (!selectedContact || selectedContact!.id !== from) ? (contact.unreadCount || 0) + 1 : contact.unreadCount
                    }
                }
                return contact
            })
            return newContacts
        })
    };

    const eventHandler = {
        [EventType.ReceiveMessage]: handleReceiveMessage,
    } as const

    const onMessage = (message: string) => {
        let payload = JSON.parse(message) as Payload<ReceiveMessage>;
        // @ts-ignore
        eventHandler[payload.event](payload);
    }


    const handleSelectContact = (contact: Contact) => {
        setSelectedContact(contact);
        // Clear unread count when selecting a contact
        setContacts(contacts.map(c =>
            c.id === contact.id ? {...c, unreadCount: 0} : c
        ));
    };

    const handleSendMessage = (content: string) => {
        chatConnection.send(JSON.stringify({
            event: EventType.SendMessage,
            data: {
                to: selectedContact!.id,
                content,
            }
        }))

        const newMessage: MessageBoxItem = {
            senderID: 0, // Current user
            senderAvatar: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png",
            senderName: "Me",
            content,
        };

        setChatHistory(prev => ({
            ...prev,
            [selectedContact!.id]: [...(prev[selectedContact!.id] || []), newMessage]
        }));

        // Update last message in contacts
        setContacts(contacts.map(c =>
            c.id === selectedContact!.id ? {...c, lastMessage: content} : c
        ));
    };

    return (

        <div className={className}>
            {contextHolder}
            <ContactList
                contacts={contacts}
                selectedContactId={selectedContact?.id}
                onSelectContact={handleSelectContact}
            />
            <div className="chat-area">
                {selectedContact ? (
                    <>
                        <div className="chat-header">
                            {selectedContact.name}
                        </div>
                        <div className="chat-messages">
                            <Messagebox
                                items={chatHistory[selectedContact.id] || []}
                            />
                        </div>
                        <ChatInput onSendMessage={handleSendMessage}/>
                    </>
                ) : (
                    <div className="no-chat-selected">
                        请选择一个联系人开始聊天
                    </div>
                )}
            </div>
        </div>
    );
};

export default styled(Index)`
    display: flex;
    width: 1280px;
    height: 860px;
    background-color: #f0f2f5;
    margin: 50px auto 0;

    .chat-area {
        flex: 1;
        display: flex;
        flex-direction: column;
        background-color: white;
    }

    .chat-header {
        padding: 16px;
        font-size: 16px;
        font-weight: 500;
        border-bottom: 1px solid #e8e8e8;
        background-color: white;
    }

    .chat-messages {
        flex: 1;
        overflow-y: auto;
        padding: 20px;
    }

    .no-chat-selected {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: center;
        color: #999;
        font-size: 16px;
    }
`;