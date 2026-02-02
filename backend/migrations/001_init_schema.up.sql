-- Создание таблицы applications
CREATE TABLE IF NOT EXISTS applications (
    id VARCHAR(36) PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    address TEXT NOT NULL,
    district VARCHAR(50) NOT NULL,
    container_id VARCHAR(36),
    service_type VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'new',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы containers
CREATE TABLE IF NOT EXISTS containers (
    id VARCHAR(36) PRIMARY KEY,
    volume INTEGER NOT NULL UNIQUE,
    price INTEGER NOT NULL,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true
);

-- Создание таблицы benefits
CREATE TABLE IF NOT EXISTS benefits (
    id VARCHAR(36) PRIMARY KEY,
    category VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    discount INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true
);

-- Создание таблицы tariffs
CREATE TABLE IF NOT EXISTS tariffs (
    id VARCHAR(36) PRIMARY KEY,
    price DECIMAL(10, 2) NOT NULL,
    valid_from TIMESTAMP NOT NULL,
    valid_to TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT true,
    description TEXT
);

-- Индексы
CREATE INDEX idx_applications_status ON applications(status);
CREATE INDEX idx_applications_created_at ON applications(created_at);
CREATE INDEX idx_tariffs_valid_from ON tariffs(valid_from);
CREATE INDEX idx_tariffs_is_active ON tariffs(is_active);

-- Вставка начальных данных для контейнеров
INSERT INTO containers (id, volume, price, description, is_active) VALUES
    ('container-120', 120, 3000, 'Если в доме живут 4-5 человек, то 120 литров окажется вполне достаточно', true),
    ('container-240', 240, 4000, 'Такая вместительность бака будет приемлема, если в вашей семье от 5 до 10 человек', true),
    ('container-1100', 1100, 19000, 'Подходит для установки во дворах офисов, магазинов или образовательных учреждений', true)
ON CONFLICT (volume) DO NOTHING;

-- Вставка начальных данных для льгот
INSERT INTO benefits (id, category, description, discount, is_active) VALUES
    ('benefit-large-family', 'large_family', 'Многодетная семья', 30, true),
    ('benefit-disabled', 'disabled', 'Инвалиды', 30, true),
    ('benefit-repressed', 'repressed', 'Репрессированные', 50, true),
    ('benefit-afghan', 'afghan_veteran', 'Ветераны Афганской войны', 100, true),
    ('benefit-wwii', 'wwii_veteran', 'Ветераны Великой Отечественной войны', 100, true)
ON CONFLICT (category) DO NOTHING;

-- Вставка начальных данных для тарифов
INSERT INTO tariffs (id, price, valid_from, valid_to, is_active, description) VALUES
    ('tariff-2024', 447.28, '2024-01-01 00:00:00', '2024-12-31 23:59:59', true, 'Тариф на 2024 год'),
    ('tariff-2025', 454.02, '2025-01-01 00:00:00', NULL, true, 'Тариф на 2025 год')
ON CONFLICT DO NOTHING;

