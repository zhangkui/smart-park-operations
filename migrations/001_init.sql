CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE IF NOT EXISTS audit_logs(id uuid PRIMARY KEY DEFAULT gen_random_uuid(), actor_id text NOT NULL, action text NOT NULL, resource text NOT NULL, resource_id text NOT NULL, ip inet, created_at timestamptz NOT NULL DEFAULT now());
CREATE TABLE IF NOT EXISTS parks(id uuid PRIMARY KEY DEFAULT gen_random_uuid(), name text NOT NULL, address text NOT NULL, status text NOT NULL DEFAULT 'active', created_at timestamptz NOT NULL DEFAULT now());
CREATE TABLE IF NOT EXISTS tenants(id uuid PRIMARY KEY DEFAULT gen_random_uuid(), name text NOT NULL, contact text NOT NULL, phone text NOT NULL, lease_status text NOT NULL DEFAULT 'active');
CREATE TABLE IF NOT EXISTS work_orders(id uuid PRIMARY KEY DEFAULT gen_random_uuid(), title text NOT NULL, description text NOT NULL, status text NOT NULL DEFAULT 'open', priority text NOT NULL DEFAULT 'normal', created_at timestamptz NOT NULL DEFAULT now());
