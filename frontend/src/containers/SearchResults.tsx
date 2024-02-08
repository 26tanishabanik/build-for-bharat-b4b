import { Container, Heading, VStack, Box, Flex, Spinner } from '@chakra-ui/react';
import SearchResultItem from '../components/SearchResultItem';
import useQuery from '../hooks/useQuery';
import { CartItemAttributes } from '../types/types';
import { useEffect, useState } from 'react';
import { SearchProductGetUUID, SearchProductPollResults } from '../apis/SearchProduct';
import { useDispatch } from 'react-redux';

const SearchResults = () => {
	const query = useQuery();
	const productName = query.get("q");
	const dispatch = useDispatch();

	const [searchResultItems, setSearchResultItems] = useState<CartItemAttributes[]>([]);
	const [isSearchComplete, setIsSearchComplete] = useState<boolean>(false);
	const [uuid, setUUID] = useState<string>("");

	useEffect(() => {
		let timeoutID: NodeJS.Timeout | null = null;

		const pollForResults = async () => {
			let attempts = 0;

			if (productName !== null && uuid !== "") {
				while (attempts < 5) {
					const searchResult = await SearchProductPollResults(productName, uuid, dispatch);
					if (searchResult.isResult) {
						setSearchResultItems(searchResult.products);
						setIsSearchComplete(true);
						setUUID("");
						break;
					} else {
						await new Promise(resolve => {
							timeoutID = setTimeout(resolve, 1000)
						});
						attempts++;
					}
				}
			}
		}

		pollForResults();

		return () => {
			if (timeoutID !== null) {
				clearTimeout(timeoutID);
			}
		}
	}, [uuid])

	useEffect(() => {
		setIsSearchComplete(false);
		setSearchResultItems([]);

		let timeoutID: NodeJS.Timeout | null = null;

		const getUUID = async () => {
			if (productName !== null) {
				const id = await SearchProductGetUUID(productName, dispatch)
				await new Promise(resolve => {
					timeoutID = setTimeout(resolve, 1000)
				});
				setUUID(id);
			}
		}

		getUUID();

		return () => {
			if (timeoutID !== null) {
				clearTimeout(timeoutID);
			}
		}
	}, [productName])

	return (
		<Container maxW='container.lg' my={4}>
			<VStack spacing={5}>
				{
					isSearchComplete ?
						(
							searchResultItems.length !== 0 ?
								searchResultItems.map((item): any =>
									<SearchResultItem
										key={item.productID}
										productID={item.productID}
										productName={item.productName}
										image={item.image}
										seller={item.seller}
										price={item.price}
										quantity={item.quantity}
									/>
								)
							:
							<Flex height="77vh" align="center" justify="center">
								<Box>
									<Heading>No Results Found</Heading>
								</Box>
							</Flex>
						)
						:
						(
							<Flex height="77vh" align="center" justify="center">
								<Box>
									<Spinner size='xl' thickness='3px' />
								</Box>
							</Flex>
						)
				}
			</VStack>
		</Container>
	);
};

export default SearchResults;