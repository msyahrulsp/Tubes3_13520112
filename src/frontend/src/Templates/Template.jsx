import React from 'react';
import { motion } from "framer-motion"

import { Navbar } from './Navbar';

export const Template = (props) => (
    <div>
        <Navbar />
        <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit= {{ opacity: 0 }}
        >
            {props.children}
        </motion.div>
    </div>
);