import { createSlice } from '@reduxjs/toolkit';

export const errorSlice = createSlice({
  name: 'error',
  initialState: {
    error: ""
  },
  reducers: {
    setError: (state, param: {type: string, payload: string}) => {
      const { payload } = param;
      state.error = payload;
    },
    removeError: (state) => {
      state.error = "";
    }
  }
});

export const { setError, removeError } = errorSlice.actions;

export default errorSlice.reducer;
