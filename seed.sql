INSERT INTO parks(name,address,status) VALUES ('智慧园区一号园','上海市浦东新区','active') ON CONFLICT DO NOTHING;
INSERT INTO tenants(name,contact,phone,lease_status) VALUES ('示范科技有限公司','管理员','13800000000','active') ON CONFLICT DO NOTHING;
