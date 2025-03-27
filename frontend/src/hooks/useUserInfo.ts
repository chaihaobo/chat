import {LoginResponse} from "../api/user.ts";
import {useLocalStorage} from "react-use";

const useUserInfo = () => {
    let [userInfo, setUserInfo] = useLocalStorage<LoginResponse | null>("user_info", null);
    return {userInfo, setUserInfo}
}

export default useUserInfo