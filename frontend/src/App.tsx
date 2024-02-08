import { Box } from '@chakra-ui/react';
import Home from './containers/Home';
import SearchResults from './containers/SearchResults';
import Header from './components/Header';
import Cart from './containers/Cart';
import { Routes, Route } from "react-router-dom";
import { useSelector } from 'react-redux';
import { IRootState } from './redux/store';

import './App.css';

export default function App() {
	const setOverlay = useSelector((state: IRootState) => state.searchInputActive.setOverlay);

	return (
    <>
		 	{/* Overlay on search from navbar */}
			<Box className={setOverlay ? 'fadeIn' : 'fadeOut'}
				w="100%" h="100%" background="rgba(0, 0, 0, 0.3)"
				position="fixed" zIndex={setOverlay ? 2 : -1}
			/>
			<Box>
				<Header />
				<Routes>
					<Route path="/" element={ <Home/> } />
					<Route path="/search" element={ <SearchResults/> } />
          <Route path="/cart" element={ <Cart/> } />
				</Routes>
			</Box>
    </>
  );
}
