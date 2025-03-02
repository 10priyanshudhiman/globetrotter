import React,{useEffect} from "react";
import { getRequestToAPI } from "../services/api";
import URLS from "../urls";
import store from "../store/store";
import { updateUser } from "../store/actions";

const FetchUserDetails = () => {

    const fetchUser = async () => {
        const userId = localStorage.getItem("userId")
        if (!userId) return 
        try {
            const url = `${URLS.login}?userId=${userId}`;
            const response = await getRequestToAPI(url);
      
            if (response.status === 200 && response.data) {
              const userData = response.data.data; 
              store.dispatch(updateUser(userData))
            }
          } catch (err) {
            console.log("Failed to fetch user details. Please try again.",err);
          } 
    }

    useEffect(() => {
        fetchUser()
    },[])
}

export default FetchUserDetails