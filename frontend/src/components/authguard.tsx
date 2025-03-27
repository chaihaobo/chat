import {FC, Fragment, ReactNode} from "react";
import {Navigate} from "react-router";
import useUserInfo from "../hooks/useUserInfo.ts";

const AuthGuard: FC<{
    children: ReactNode
}> = ({children}) => {
    let {userInfo} = useUserInfo();
    if (!userInfo?.id) {
        return <Navigate to="/login" replace/>;
    }
    return (
        <Fragment>
            {children}
        </Fragment>
    )
}


export default AuthGuard