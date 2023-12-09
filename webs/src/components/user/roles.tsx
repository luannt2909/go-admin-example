import React from 'react'

export const RoleAdmin = 1
export const RoleGuest = 2

export const Roles = [
    {id: RoleAdmin, name: "Admin"},
    {id: RoleGuest, name: "Guest"},
];
export const GetRole = function (role) {
    switch (role) {
        case RoleAdmin:
            return "Admin"
        case RoleGuest:
            return "Guest"
        default:
            return "Unknown"
    }
}
