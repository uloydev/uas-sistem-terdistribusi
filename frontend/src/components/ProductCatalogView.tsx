import axios from 'axios'
import { useRouter } from 'next/router'
import React from 'react'
import { BASE_URL_PRODUCT_API } from '../../utils/const'
import { currencyIDR } from '../../utils/helpers'
import ActionTableButton from './ActionTableButton'
import CustomTable from './CustomTable'

import Template from './Template'
import WrapperCard from './WrapperCard'

interface ProductProp {
    description: string,
    id: number,
    price: number,
    stock: number,
    title: string,
    created_at: number,
}

const ProductCatalogView = () => {
    const router = useRouter();
    const [product, setProduct] = React.useState<ProductProp[] | []>([])
    const [isLoading, setIsLoading] = React.useState<boolean>(true)

    const handleDeleteProduct = async (id: number) => {
        try {
            const response = await axios.delete(BASE_URL_PRODUCT_API + `/${id}?master=supersecret`)
            fetchDataProduct()
        } catch (err) {
            console.log(err)
        }
    }

    const fetchDataProduct = async () => {
        try {
            const { data: result } = await axios.get(BASE_URL_PRODUCT_API + "?master=supersecret")
            const { data } = result
            if (data === null) {
                setProduct([])
                setIsLoading(false)
                return
            }
            const filtererData = data.map((item: any) => {
                return {
                    ...item,
                    price: currencyIDR(item.price),
                    action: <ActionTableButton onUpdate={() => router.push(`/product/update/${item.id}`)} onDelete={() => handleDeleteProduct(item.id)} />
                }
            })
            setProduct([...filtererData])
            setIsLoading(false)
        } catch (err) {
            setProduct([])
            setIsLoading(false)
            console.log(err)
        }
    }
    const fetcher = React.useRef(fetchDataProduct);

    const columns = [
        {
            name: 'Id',
            selector: (row: any) => row.id
        },
        {
            name: 'Title',
            selector: (row: any) => row.title
        },
        {
            name: 'Stock',
            selector: (row: any) => row.stock
        },
        {
            name: 'Price',
            selector: (row: any) => row.price
        },
        {
            name: 'Description',
            selector: (row: any) => row.description
        },
        {
            name: 'Action',
            selector: (row: any) => row.action
        }
    ]

    React.useEffect(() => {
        fetcher.current()
    }, [fetcher])

    return (
        <Template title="Product Catalog">
            <WrapperCard>
                {isLoading ? (<div>loading</div>) : (<CustomTable columns={columns} data={product} />)}
            </WrapperCard>
        </Template>
    )
}

export default ProductCatalogView