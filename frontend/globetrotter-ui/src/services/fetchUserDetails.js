import { useEffect } from "react";
import { getRequestToAPI } from "../services/api";
import URLS from "../urls";
import store from "../store/store";
import { updateUser } from "../store/actions";

const FetchUserDetails = () => {

  useEffect(() => {
    // Create an abort controller to cancel the fetch request if needed
    const controller = new AbortController();
    const { signal } = controller;

    const fetchUser = async () => {
      const userId = localStorage.getItem("userId");
      if (!userId) return;

      try {
        const url = `${URLS.login}?userId=${userId}`;
        const response = await getRequestToAPI(url, { signal }); // Pass the signal to the fetch call

        if (response.status === 200 && response.data) {
          const userData = response.data.data;
          store.dispatch(updateUser(userData)); // Dispatch the action to update user details
        }
      } catch (err) {
        // Check if the error is due to the request being aborted
        if (err.name !== "AbortError") {
          console.log("Failed to fetch user details. Please try again.", err);
        }
      }
    };

    // Call fetchUser function
    fetchUser();

    // Cleanup function to abort fetch request if component unmounts
    return () => {
      controller.abort(); // Abort the fetch when component unmounts
    };
  }, []); // Empty dependency array ensures it runs only once when the component mounts

  return null; // This component doesn't render anything, it's just for side effects
};

export default FetchUserDetails;
