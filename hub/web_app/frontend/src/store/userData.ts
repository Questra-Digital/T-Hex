import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { UserData } from "@/types/user";

interface UserDataState {
    userData: UserData;
}

const initialState: UserDataState = {
    userData: {
        user_id: 0,
    },
}

const userDataSlice = createSlice({
    name: 'userData',
    initialState,
    reducers: {
        setUserData: (state, action: PayloadAction<UserData>) => {
            state.userData = action.payload;
        },
        updateUserData: (state, action: PayloadAction<UserData>) => {
            state.userData = { ...state.userData, ...action.payload };
        }
        ,

    },
});

export const { setUserData, updateUserData } = userDataSlice.actions;
export default userDataSlice.reducer;

// Selectors
export const selectUserData = (state: { userData: UserDataState }) => state.userData.userData;