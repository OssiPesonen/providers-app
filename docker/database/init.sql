-- providers table
-- id should be bigserial
CREATE TABLE 
"public"."providers" (
    "id" bigserial not null,
    "name" varchar(255) not null,
    "city" varchar(255) not null,
    "region" varchar(255) not null,
    "line_of_business" varchar(255) not null,
    "keywords" varchar(255) not null,
    "userId" bigserial not null,
    constraint "providers_pkey" primary key ("id")
);

insert into "public"."providers" ("city", "keywords", "line_of_business", "name", "region", "userId") values 
('Helsinki', 'personal,training,gym,wellness,health', 'Health and Wellness', 'Healthy Personal Training', 'Uusimaa', 1), 
('Stockholm', 'yoga,wellness,lifestyle,health', 'Health and Wellness', 'Healthy Yoga Instructor', 'Stockholm', 1), 
('Oslo', 'agile,coaching,coach,software,development,teams', 'Information technology', 'Healthy Agile Coaching', 'Akershus', 1);

-- end of providers table

-- users table 

CREATE TABLE
  "public"."users" (
    "id" bigserial not null,
    "username" varchar(255) not null default NOW(),
    "email" varchar(255) not null,
    "password" text not null,
    "salt" varchar(255) not null,
    constraint "user_pkey" primary key ("id")
  );
  
  -- end user table

CREATE TABLE
  "public"."sessions" (
    "token" varchar(255) not null,
    "expires" timestamp without time zone not null default now(),
    "user_id" integer not null,
    constraint "sessions_pkey" primary key ("token","user_id")
);

-- fts for providers
ALTER TABLE providers ADD COLUMN search_vector tsvector generated always as (
  to_tsvector('english', 
              coalesce(name, '') || ' ' || 
              coalesce(city, '') || ' ' ||
              coalesce(region, '') || ' ' ||
              coalesce(line_of_business, '')
             )
) stored;

CREATE INDEX textsearch_idx ON "providers" USING GIN ("search_vector");

/*

Here's some dummy data to seed providers for search

INSERT INTO providers (name, city, region, line_of_business, user_id) VALUES
('TechNova Solutions', 'New York', 'New York', 'Technology', 1),
('GreenLeaf Organics', 'Los Angeles', 'California', 'Agriculture', 1),
('SkyHigh Airways', 'Chicago', 'Illinois', 'Aviation', 1),
('Quantum Motors', 'Detroit', 'Michigan', 'Automotive', 1),
('Oceanic Ventures', 'Miami', 'Florida', 'Marine', 1),
('EcoBuild Constructions', 'Austin', 'Texas', 'Construction', 1),
('FutureEdge AI', 'San Francisco', 'California', 'Artificial Intelligence', 1),
('HealthHive Pharma', 'Boston', 'Massachusetts', 'Pharmaceuticals', 1),
('GlobeTrotter Travel', 'Seattle', 'Washington', 'Travel', 1),
('FreshFields Agro', 'Portland', 'Oregon', 'Agriculture', 1),
('UrbanSmart Homes', 'Atlanta', 'Georgia', 'Real Estate', 1),
('FinTrust Advisory', 'Dallas', 'Texas', 'Finance', 1),
('BrightPath Education', 'Philadelphia', 'Pennsylvania', 'Education', 1),
('NovaTech Labs', 'Denver', 'Colorado', 'Technology', 1),
('SolarFlare Energy', 'Phoenix', 'Arizona', 'Renewable Energy', 1),
('PrimeFleet Logistics', 'Houston', 'Texas', 'Logistics', 1),
('Skyline Architects', 'Las Vegas', 'Nevada', 'Architecture', 1),
( 'AquaPure Systems', 'San Diego', 'California', 'Water Technology', 1),
('SafeShield Security', 'Orlando', 'Florida', 'Security', 1),
('CloudCore IT', 'San Jose', 'California', 'Information Technology', 1),
('BrightFuture Biotech', 'Raleigh', 'North Carolina', 'Biotechnology', 1),
('EverGreen Forestry', 'Anchorage', 'Alaska', 'Forestry', 1),
('NextGen Robotics', 'Pittsburgh', 'Pennsylvania', 'Robotics', 1),
('AlphaComms Media', 'Charlotte', 'North Carolina', 'Media', 1),
('Zenith Manufacturing', 'Milwaukee', 'Wisconsin', 'Manufacturing', 1),
('GlobalConnect Telecom', 'Minneapolis', 'Minnesota', 'Telecommunications', 1),
('EcoTrend Fashion', 'Nashville', 'Tennessee', 'Fashion', 1),
('PureHarvest Foods', 'Salt Lake City', 'Utah', 'Food Production', 1),
('BlueWave Marine', 'Tampa', 'Florida', 'Marine', 1),
('Infinity Entertainment', 'Las Vegas', 'Nevada', 'Entertainment', 1),
('CoreMetrics Analytics', 'Cleveland', 'Ohio', 'Analytics', 1),
('RapidFix Auto', 'Kansas City', 'Missouri', 'Automotive', 1),
('GreenGrid Power', 'Columbus', 'Ohio', 'Renewable Energy', 1),
('AeroDynamics Inc.', 'Tulsa', 'Oklahoma', 'Aerospace', 1),
('UrbanPulse Media', 'Indianapolis', 'Indiana', 'Media', 1),
('RiverBend Resorts', 'Sacramento', 'California', 'Hospitality', 1),
('ProHealth Clinics', 'Madison', 'Wisconsin', 'Healthcare', 1),
('Skybridge Logistics', 'Oklahoma City', 'Oklahoma', 'Logistics', 1),
('VividVision Studios', 'Memphis', 'Tennessee', 'Entertainment', 1),
('TrailBlaze Outdoor', 'Boise', 'Idaho', 'Outdoor Equipment', 1),
('PolarTech Systems', 'Juneau', 'Alaska', 'Technology', 1),
('SmartUrban Designs', 'Richmond', 'Virginia', 'Architecture', 1),
('HarvestTime Grocers', 'Hartford', 'Connecticut', 'Retail', 1),
('AlphaEdge AI', 'Boulder', 'Colorado', 'Artificial Intelligence', 1),
('DeepBlue Tech', 'Honolulu', 'Hawaii', 'Marine Technology', 1),
( 'SwiftLink Logistics', 'Louisville', 'Kentucky', 'Logistics', 1),
( 'PeakPerformance Sports', 'Denver', 'Colorado', 'Sports', 1),
( 'ClearView Optics', 'Albany', 'New York', 'Optics', 1),
( 'SolarSphere Energy', 'Santa Fe', 'New Mexico', 'Renewable Energy', 1),
( 'VibrantMinds Education', 'Providence', 'Rhode Island', 'Education', 1);

*/