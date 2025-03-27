import {FC} from "react";
import {Avatar, Button, List} from "antd";
import {LogoutOutlined} from '@ant-design/icons';
import styled from "styled-components";
import useUserInfo from "../../../hooks/useUserInfo";

export interface Contact {
    id: number;
    name: string;
    avatar: string;
    lastMessage?: string;
    unreadCount?: number;
}

interface ContactListProps {
    className?: string;
    contacts: Contact[];
    selectedContactId?: number;
    onSelectContact: (contact: Contact) => void;
}

const ContactList: FC<ContactListProps> = ({
                                               className,
                                               contacts,
                                               selectedContactId,
                                               onSelectContact
                                           }) => {
    return (
        <div className={className}>
            <div className="header">
                <div className="title">联系人列表</div>
                <div className="user-info">
                    <Avatar src={useUserInfo().userInfo?.avatar}/>
                    <span className="username">{useUserInfo().userInfo?.username}</span>
                    <Button
                        type="text"
                        icon={<LogoutOutlined/>}
                        onClick={() => {
                            localStorage.removeItem('user_info');
                            window.location.href = '/login';
                        }}
                    />
                </div>
            </div>
            <List
                dataSource={contacts}
                renderItem={(contact) => (
                    <List.Item
                        className={`contact-item ${contact.id === selectedContactId ? 'selected' : ''}`}
                        onClick={() => onSelectContact(contact)}
                    >
                        <List.Item.Meta
                            avatar={<Avatar src={contact.avatar}/>}
                            title={contact.name}
                            description={contact.lastMessage}
                        />
                        {contact.unreadCount ? (
                            <div className="unread-badge">{contact.unreadCount}</div>
                        ) : null}
                    </List.Item>
                )}
            />
        </div>
    );
};

export default styled(ContactList)`
    width: 280px;
    height: 100%;
    background-color: white;
    border-right: 1px solid #e8e8e8;

    .header {
        padding: 16px;
        font-size: 16px;
        font-weight: 500;
        border-bottom: 1px solid #e8e8e8;
        display: flex;
        justify-content: space-between;
        align-items: center;

        .user-info {
            display: flex;
            align-items: center;
            gap: 8px;
        }
    }

    .contact-item {
        padding: 12px 16px;
        cursor: pointer;
        transition: background-color 0.3s;

        &:hover {
            background-color: #f5f5f5;
        }

        &.selected {
            background-color: #e6f7ff;
        }
    }

    .unread-badge {
        background-color: #ff4d4f;
        color: white;
        border-radius: 10px;
        padding: 0 6px;
        font-size: 12px;
        min-width: 20px;
        height: 20px;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .ant-list-item-meta-description {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        max-width: 200px;
    }
`;