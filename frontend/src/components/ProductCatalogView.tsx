import axios from 'axios'
import React from 'react'
import { BASE_URL_PRODUCT_API } from '../../utils/const'
import { currencyIDR } from '../../utils/helpers'
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
    const [product, setProduct] = React.useState<ProductProp[] | []>([])
    const [isLoading, setIsLoading] = React.useState<boolean>(true)

    const columns = [
        {
            name: 'id',
            selector: (row: any) => row.id
        },
        {
            name: 'title',
            selector: (row: any) => row.title
        },
        {
            name: 'stock',
            selector: (row: any) => row.stock
        },
        {
            name: 'price',
            selector: (row: any) => row.price
        },
        {
            name: 'description',
            selector: (row: any) => row.description
        }
    ]

    React.useEffect(() => {
        const fetchDataProduct = async () => {
            try {
                const { data: result } = await axios.get(BASE_URL_PRODUCT_API + "?master=supersecret")
                const { data } = result
                let filtererData = data.map((item: any) => {
                    return { ...item, price: currencyIDR(item.price) }
                })
                setProduct([...filtererData])
                setIsLoading(false)
            } catch (err) {
                console.log(err)
            }
        }
        fetchDataProduct()
    }, [])

    return (
        <Template title="Product Catalog">
            <WrapperCard>
                {isLoading ? (<div>loading</div>) : (<CustomTable columns={columns} data={product} />)}
            </WrapperCard>
        </Template>
    )
}

export default ProductCatalogView