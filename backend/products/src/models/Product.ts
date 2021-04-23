import { Entity, Column, PrimaryColumn } from "typeorm";
import { Discount } from "./Discount";

@Entity('product')
export class Product {

    @PrimaryColumn('uuid')
    id!: string;

    @Column('int')
    price_in_cents!: number;

    @Column('text')
    title!: string;

    @Column('text')
    description!: string;

    discount: Discount = { percentage: 0, value_in_cents: 0}

    setDiscount(percentage: number, value_in_cents: number){
        this.discount = {
            percentage,
            value_in_cents
        };
    }
}

