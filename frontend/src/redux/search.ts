import { createSlice } from '@reduxjs/toolkit';

export const backgroundOverlayForSearchSlice = createSlice({
  name: 'backgroundOverlayForSearch',
  initialState: {
    setOverlay: false
  },
  reducers: {
    setBackgroundOverlay: (state, param: {type: string, payload: string}) => {
      const { payload } = param;
      if (payload === "/") {
        return;
      }
      state.setOverlay = true;
    },
    removeBackgroundOverlay: (state, param: {type: string, payload: string}) => {
      const { payload } = param;
      if (payload === "/") {
        return;
      }
      state.setOverlay = false;
    }
  }
});

export const { setBackgroundOverlay, removeBackgroundOverlay } = backgroundOverlayForSearchSlice.actions;

export default backgroundOverlayForSearchSlice.reducer;
