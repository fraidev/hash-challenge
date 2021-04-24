import { ConnectionOptions } from "typeorm";
import { Product } from "../models/Product";

export const TYPEORM_OPTIONS: ConnectionOptions = {
    type: "postgres",
    host: process.env.NODE_ENV === "production" ? "hash-db" : "localhost",
    port: 5432,
    username: "postgres",
    password: "YOUR_PASSWORD_HERE",
    database: "hash",
    entities: [Product]
}