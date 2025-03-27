import {FC, useEffect, useState} from 'react';
import {useNavigate} from 'react-router';
import {message, Spin} from 'antd';
import {login} from '../../api/user';
import useUserInfo from "../../hooks/useUserInfo.ts";

const Callback: FC = () => {
    const navigate = useNavigate();
    const [messageApi, contextHolder] = message.useMessage();
    const [loading, setLoading] = useState(true);
    const code = new URLSearchParams(window.location.search).get('code');
    // @ts-ignore
    let {userInfo, setUserInfo} = useUserInfo()
    if (!code) {
        messageApi.error('授权失败，请重试');
        navigate('/login');
        return;
    }

    useEffect(() => {
        const handleLogin = async () => {
            try {
                const userInfo = await login(code);
                setUserInfo(userInfo)
                messageApi.success('登录成功');
                setTimeout(() => {
                    navigate('/chat');
                }, 1000)
            } catch (error) {
                console.error('Login failed:', error);
                messageApi.error('登录失败，请重试');
                setTimeout(() => {
                    navigate('/login');
                }, 1000)

            } finally {
                setLoading(false);
            }
        };
        handleLogin();
    }, []);

    return (
        <div style={{display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh'}}>
            {contextHolder}
            <Spin spinning={loading} tip="正在处理GitHub授权..."/>
        </div>
    );
};

export default Callback;