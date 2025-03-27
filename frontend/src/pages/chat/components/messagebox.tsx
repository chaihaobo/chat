import {FC} from "react";
import {Avatar} from "antd";
import styled from "styled-components";
import useUserInfo from "../../../hooks/useUserInfo.ts";

export interface MessageBoxProps {
    items: MessageBoxItem[]
    className?: string
}

export interface MessageBoxItem {
    senderID: number
    senderAvatar: string
    senderName: string
    content: string
}

const Messagebox: FC<MessageBoxProps> = ({items, className}) => {
    let {userInfo} = useUserInfo();
    return (
        <div className={className}>
            {items.map((item, index) => (
                <div
                    key={index}
                    className={`message-item ${item.senderID === userInfo?.id ? 'message-self' : ''}`}
                >
                    <Avatar src={item.senderAvatar} className="avatar"/>
                    <div className="message-content">
                        <div className="sender-name">{item.senderName}</div>
                        <div className="message-bubble">{item.content}</div>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default styled(Messagebox)`
    display: flex;
    flex-direction: column;
    gap: 16px;

    .message-item {
        display: flex;
        gap: 12px;
        align-items: flex-start;

        &.message-self {
            flex-direction: row-reverse;

            .message-content {
                align-items: flex-end;
            }

            .message-bubble {
                background-color: #1890ff;
                color: white;
                border-radius: 8px 2px 8px 8px;
            }

            .sender-name {
                text-align: right;
            }
        }
    }

    .message-content {
        display: flex;
        flex-direction: column;
        gap: 4px;
    }

    .sender-name {
        font-size: 12px;
        color: #999;
    }

    .message-bubble {
        background-color: #f5f5f5;
        padding: 8px 12px;
        border-radius: 2px 8px 8px 8px;
        max-width: 560px;
        word-wrap: break-word;
    }
`;