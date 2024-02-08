import { Search2Icon } from '@chakra-ui/icons';
import { InputRightElement, Input, InputGroup, IconButton, VStack, Box, Center } from '@chakra-ui/react';

import debounce from 'lodash.debounce';
import { useState, useMemo, useEffect, useRef } from 'react';

import { Link, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { setBackgroundOverlay, removeBackgroundOverlay } from '../redux/search';
import { useLocation } from 'react-router-dom';
import SearchAutocomplete from '../apis/SearchAutocomplete';
import useQuery from '../hooks/useQuery';

function Search() {
	const query = useQuery();
	const productName = query.get("q");

	const [searchQuery, setSearchQuery] = useState<string>(productName ?? "");
	const [inputFocused, setInputFocused] = useState<boolean>(false);
	const [mouseInSearchResults, setMouseInSearchResults] = useState<boolean>(false);
	const [listToDisplay, setListToDisplay] = useState<string[]>([]);

	const dispatch = useDispatch();
	const location = useLocation();
	const navigate = useNavigate();
	const inputRef = useRef<HTMLInputElement|null>(null);

	useEffect(() => {
		if (searchQuery !== "") {
			SearchAutocomplete(searchQuery, dispatch)
			.then((result: string[]) => {
				setListToDisplay(result);
			});
		}
	}, [searchQuery]);

	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchQuery(e.target.value);
  };

	const debouncedResults = useMemo(() => {
    return debounce(handleChange, 300);
  }, []);

  useEffect(() => {
    return () => {
			debouncedResults.cancel();
		};
  });

	return (
		<>
			<Center>
				<Box position="relative">
					<InputGroup w="35vw" bgColor="white" size={location.pathname === '/'? 'lg' : 'md'}>
						<Input
							ref={inputRef}
							size={location.pathname === '/'? 'lg' : 'md'}
							value={searchQuery} focusBorderColor='black' placeholder='Search'
							onFocus={() => {setInputFocused(true); dispatch(setBackgroundOverlay(location.pathname))}}
							onChange={e => handleChange(e)}
							onBlur={() => {setInputFocused(false); dispatch(removeBackgroundOverlay(location.pathname))}} 
							onKeyDown={(e) => {
								if (e.key === 'Enter' && searchQuery !== "") {
									navigate(`/search?q=${searchQuery.toLowerCase()}`);
									setInputFocused(false);
									if (inputRef.current !== null) {
										inputRef.current.blur();
									}
									dispatch(removeBackgroundOverlay(location.pathname));
								}
							}} />
						<InputRightElement>
							<Link to={`/search?q=${searchQuery.toLowerCase()}`}>
								<IconButton
									aria-label='Search'
									colorScheme='transparent'
									size={location.pathname === '/'? 'lg' : 'md'}
									color='grey'
									icon={<Search2Icon />}
									isDisabled={searchQuery === ""}
								/>
							</Link>
						</InputRightElement>
					</InputGroup>
					<Box onMouseEnter={() => setMouseInSearchResults(true)} onMouseLeave={() => setMouseInSearchResults(false)}
						visibility={searchQuery !== "" && listToDisplay !== null && listToDisplay.length !== 0 && (inputFocused || mouseInSearchResults) ? 'visible' : 'hidden'}
						mt={1} p={1.5} zIndex={1} borderRadius={10} w="35vw" position="absolute" backgroundColor="white">
						<VStack spacing={2} align='stretch'>
							{
								listToDisplay?.map((fruit, i) => {
									return (
										<Link key={i} to={`/search?q=${fruit}`} onClick={() => setSearchQuery("")}>
											<Box key={i} p={2} rounded="md" bgColor="lightgray">
												{fruit}
											</Box>
										</Link>
									);
								})
							}
						</VStack>
					</Box>
				</Box>
			</Center>
		</>
	);
};

export default Search;