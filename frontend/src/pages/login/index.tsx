import {FC} from 'react';
import {Button, Form, Input, message} from 'antd';
import {GithubOutlined, LockOutlined, UserOutlined} from '@ant-design/icons';
import {styled} from 'styled-components';
import {loginWithPassword} from '../../api/user';
import {useNavigate} from 'react-router';
import useUserInfo from '../../hooks/useUserInfo';

const GITHUB_CLIENT_ID = import.meta.env.VITE_GITHUB_CLIENT_ID || '';
const GITHUB_REDIRECT_URI = import.meta.env.VITE_GITHUB_REDIRECT_URI || 'http://localhost:5173/auth/callback';

const Login: FC<{ className?: string }> = ({className}) => {
    const navigate = useNavigate();
    const [messageApi, contextHolder] = message.useMessage();
    const {setUserInfo} = useUserInfo();
    const source = new URLSearchParams(window.location.search).get('source');

    const handleGithubLogin = () => {
        const githubAuthUrl = `https://github.com/login/oauth/authorize?client_id=${GITHUB_CLIENT_ID}&redirect_uri=${GITHUB_REDIRECT_URI + '?source=' + source}&scope=user`;
        window.location.href = githubAuthUrl;
    };

    const onFinish = async (values: { username: string; password: string }) => {
        try {
            const userInfo = await loginWithPassword(values);
            setUserInfo(userInfo);
            messageApi.success('登录成功');
            navigate(source ?? '/chat');
        } catch (error) {
            messageApi.error('登录失败，请检查用户名和密码');
        }
    };


    return (
        <div className={className}>
            {contextHolder}
            <div className="login-container">
                <h1>Welcome to Chat</h1>
                <Form
                    name="login"
                    onFinish={onFinish}
                    className="login-form"
                >
                    <Form.Item
                        name="username"
                        rules={[{required: true, message: '请输入用户名'}]}
                    >
                        <Input
                            prefix={<UserOutlined/>}
                            placeholder="用户名"
                            size="large"
                        />
                    </Form.Item>
                    <Form.Item
                        name="password"
                        rules={[{required: true, message: '请输入密码'}]}
                    >
                        <Input.Password
                            prefix={<LockOutlined/>}
                            placeholder="密码"
                            size="large"
                        />
                    </Form.Item>
                    <Form.Item>
                        <Button
                            type="primary"
                            htmlType="submit"
                            size="large"
                            block
                        >
                            登录
                        </Button>
                    </Form.Item>
                </Form>
                <div className="divider">
                    <span>或</span>
                </div>
                <Button
                    type="default"
                    icon={<GithubOutlined/>}
                    size="large"
                    onClick={handleGithubLogin}
                    block
                >
                    使用GitHub登录
                </Button>
            </div>
        </div>
    );
};

export default styled(Login)`
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background-color: #f0f2f5;

    .login-container {
        text-align: center;
        padding: 40px;
        background: white;
        border-radius: 8px;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);

        h1 {
            margin-bottom: 32px;
            color: #1890ff;
        }

        .login-form {
            margin-bottom: 24px;
        }

        .divider {
            position: relative;
            text-align: center;
            margin: 16px 0;

            &::before,
            &::after {
                content: '';
                position: absolute;
                top: 50%;
                width: 45%;
                height: 1px;
                background-color: #e8e8e8;
            }

            &::before {
                left: 0;
            }

            &::after {
                right: 0;
            }

            span {
                padding: 0 8px;
                color: #999;
                background-color: white;
            }
        }
    }
`;