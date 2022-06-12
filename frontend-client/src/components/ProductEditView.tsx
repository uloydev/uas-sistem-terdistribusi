import axios from 'axios'
import { useRouter } from 'next/router'
import React, { useId } from 'react'
import { BASE_URL_PRODUCT_API } from '../../utils/const'
import buttonStyles from '../../styles/Button.module.css'
import { config } from '../../utils/config'

import Button from './Button'
import Input from './Input'
import Label from './Label'

import Template from './Template'
import TextArea from './TextArea'
import WrapperCard from './WrapperCard'

const ProductEditView = ({ product }: any) => {
    const router = useRouter()
    const id = useId()
    const [description, setDescription] = React.useState(product.description)
    const [title, setTitle] = React.useState(product.title)
    const [price, setPrice] = React.useState<number>(product.price)
    const [stock, setStock] = React.useState<number>(product.stock)

    const requestCreateProduct = async () => {
        const payload = {
            description,
            price: Number(price),
            stock: Number(stock),
            title
        }
        try {
            const response = await axios.put(`${BASE_URL_PRODUCT_API}/${product.id}`, payload, config)
            router.push("/")
        }
        catch (err) {
            alert("something wrong in here")
            console.log(err)
        }

    }
    const handleChangeDesc = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        setDescription(e.target.value)
    }

    const handleChangeTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
        setTitle(e.target.value)
    }

    const handleChangePrice = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPrice(e.target.value as any)
    }

    const handleChangeStock = (e: React.ChangeEvent<HTMLInputElement>) => {
        setStock(e.target.value as any)
    }

    return (
        <Template title='Product Edit'>
            <section className='mt-6'>
                <WrapperCard>
                    <div >
                        <h4 className='mb-7'>Create Product</h4>
                        <div className='space-y-4'>
                            <div>
                                <Label id={`${id}-title`}>Product Title</Label>
                                <Input
                                    id={`${id}-title`}
                                    value={title}
                                    type="text"
                                    onChange={handleChangeTitle}
                                    placeholder='Product title' />
                            </div>
                            <div>
                                <Label id={`${id}-price`}>Product Price</Label>
                                <Input
                                    id={`${id}-price`}
                                    value={price}
                                    type="number"
                                    onChange={handleChangePrice}
                                    placeholder='Product Price' />
                            </div>
                            <div>
                                <Label id={`${id}-stock`}>Product Stock</Label>
                                <Input
                                    id={`${id}-stock`}
                                    value={stock}
                                    type="number"
                                    onChange={handleChangeStock}
                                    placeholder='Product Stock' />
                            </div>
                            <div>
                                <Label id={`${id}-description`}>Product Description</Label>
                                <TextArea
                                    id={`${id}-description`}
                                    value={description}
                                    onChange={handleChangeDesc}
                                />
                            </div>
                        </div>
                        <div className='my-7 w-32'>
                            <Button
                                className={`${buttonStyles['btn-lavender']} hover:-translate-y-1.5 duration-200 py-2 rounded`}
                                onClick={requestCreateProduct}
                            >
                                Edit
                            </Button>
                        </div>
                    </div>
                </WrapperCard>
            </section>
        </Template>
    )
}

export default ProductEditView