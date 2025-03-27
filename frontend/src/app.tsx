import {FC} from "react";
import Layout from "./layout";
import {BrowserRouter, Navigate, Route, Routes} from "react-router";
import Login from "./pages/login";
import Chat from "./pages/chat";
import Callback from "./pages/auth/callback.tsx";
import GlobalStyle from "./globalstyle.ts";
import {message} from "antd";
import AuthGuard from "./components/authguard.tsx";


const App: FC = () => {
    const [_, contextHolder] = message.useMessage();
    return (
        <BrowserRouter>
            {contextHolder}
            <Layout>
                <Routes>
                    <Route path="/login" element={<Login/>}/>
                    <Route path="/auth/callback" element={<Callback/>}/>
                    <Route path="/" element={<AuthGuard><Navigate to="/chat" replace/></AuthGuard>}/>
                    <Route path="/chat" element={<AuthGuard><Chat/></AuthGuard>}/>
                </Routes>
            </Layout>
            <GlobalStyle/>
        </BrowserRouter>
    );
};

export default App;