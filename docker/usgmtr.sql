-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.0
-- PostgreSQL version: 9.6
-- Project Site: pgmodeler.com.br
-- Model Author: ---


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: usgmtr | type: DATABASE --
-- -- DROP DATABASE IF EXISTS usgmtr;
-- CREATE DATABASE usgmtr
-- ;
-- -- ddl-end --
-- 

-- object: public."Psc" | type: TABLE --
-- DROP TABLE IF EXISTS public."Psc" CASCADE;
CREATE TABLE public."Psc"(
	id serial NOT NULL,
	"endPoint" character varying(256) NOT NULL,
	port integer NOT NULL DEFAULT 7444,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "Psc_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."PscHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."PscHistory" CASCADE;
CREATE TABLE public."PscHistory"(
	id serial NOT NULL,
	"pscId" integer NOT NULL,
	"endPoint" character varying(256) NOT NULL,
	port integer NOT NULL DEFAULT 7444,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "PscHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."Vc" | type: TABLE --
-- DROP TABLE IF EXISTS public."Vc" CASCADE;
CREATE TABLE public."Vc"(
	id serial NOT NULL,
	active boolean NOT NULL,
	"endPoint" character varying(256) NOT NULL,
	port integer NOT NULL,
	"userName" character varying(256) NOT NULL,
	password character varying(256) NOT NULL,
	sso integer NOT NULL DEFAULT 0,
	"pscId" integer,
	"fullName" character varying(256) NOT NULL,
	version character varying(256) NOT NULL,
	build character varying(256) NOT NULL,
	"productLineId" character varying(64) NOT NULL,
	"instanceUuid" character varying(256) NOT NULL,
	"licenseEditionKey" character varying(256) NOT NULL,
	"licenseKey" character varying(256) NOT NULL,
	"licenseName" character varying(256) NOT NULL,
	"licenseTotal" integer NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "Vc_pk" PRIMARY KEY (id),
	CONSTRAINT "VcServer_chk" CHECK ((((sso <> 2) AND ("pscId" IS NULL)) OR ((sso = 2) AND ("pscId" IS NOT NULL))))

);
-- ddl-end --

-- object: public."VcHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcHistory" CASCADE;
CREATE TABLE public."VcHistory"(
	id serial NOT NULL,
	"vcId" integer NOT NULL,
	active boolean NOT NULL,
	"endPoint" character varying(256) NOT NULL,
	port integer NOT NULL,
	"userName" character varying(256) NOT NULL,
	password character varying(256) NOT NULL,
	sso integer NOT NULL DEFAULT 0,
	"pscId" integer,
	"fullName" character varying(256) NOT NULL,
	version character varying(256) NOT NULL,
	build character varying(256) NOT NULL,
	"productLineId" character varying(64) NOT NULL,
	"instanceUuid" character varying(256) NOT NULL,
	"licenseEditionKey" character varying(256) NOT NULL,
	"licenseKey" character varying(256) NOT NULL,
	"licenseName" character varying(256) NOT NULL,
	"licenseTotal" integer NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcRegistrationHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcRegistrationHistory" CASCADE;
CREATE TABLE public."VcRegistrationHistory"(
	id serial NOT NULL,
	"vcId" integer,
	"endPoint" character varying(256) NOT NULL,
	port integer NOT NULL,
	"userName" character varying(256),
	password character varying(256),
	"pscId" integer,
	status character varying(256) NOT NULL,
	message character varying(2048),
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcRegistrationHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."User" | type: TABLE --
-- DROP TABLE IF EXISTS public."User" CASCADE;
CREATE TABLE public."User"(
	id smallserial NOT NULL,
	username character varying(128) NOT NULL,
	password character varying(128) NOT NULL,
	salt character varying(128) NOT NULL,
	active boolean NOT NULL,
	"changedByUserId" integer NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "User_pkey" PRIMARY KEY (id),
	CONSTRAINT "User_uk" UNIQUE (username)

);
-- ddl-end --

-- object: public."UserHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."UserHistory" CASCADE;
CREATE TABLE public."UserHistory"(
	id bigserial,
	"userId" smallint,
	username character varying(128) NOT NULL,
	password character varying(128) NOT NULL,
	salt character varying(128) NOT NULL,
	active boolean NOT NULL,
	"changedByUserId" integer NOT NULL,
	"changeTime" timestamp with time zone NOT NULL
);
-- ddl-end --

-- object: public."UserLoginHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."UserLoginHistory" CASCADE;
CREATE TABLE public."UserLoginHistory"(
	id bigserial NOT NULL,
	"userId" smallint,
	"loginTime" timestamp with time zone NOT NULL,
	"logoutTime" timestamp with time zone,
	"logoutReason" character varying(256),
	message character varying(2048),
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "UserLoginHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."UserPriv" | type: TABLE --
-- DROP TABLE IF EXISTS public."UserPriv" CASCADE;
CREATE TABLE public."UserPriv"(
	id serial NOT NULL,
	"userId" smallint NOT NULL,
	"roleId" smallint NOT NULL,
	"changedByUserId" integer NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "UserPriv_pk" PRIMARY KEY (id),
	CONSTRAINT "UserPriv_uk" UNIQUE ("userId","roleId")

);
-- ddl-end --

-- object: public."UserPrivHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."UserPrivHistory" CASCADE;
CREATE TABLE public."UserPrivHistory"(
	id serial NOT NULL,
	"userPrivId" integer NOT NULL,
	"userId" smallint NOT NULL,
	"roleId" smallint,
	"changedByUserId" integer,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "UserPrivHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."Role" | type: TABLE --
-- DROP TABLE IF EXISTS public."Role" CASCADE;
CREATE TABLE public."Role"(
	id smallserial NOT NULL,
	"roleName" varchar(64) NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "Role_pk" PRIMARY KEY (id),
	CONSTRAINT "Role_uk" UNIQUE ("roleName")

);
-- ddl-end --

-- object: public."ServiceProvider" | type: TABLE --
-- DROP TABLE IF EXISTS public."ServiceProvider" CASCADE;
CREATE TABLE public."ServiceProvider"(
	id serial NOT NULL,
	"partnerName" character varying(256) NOT NULL,
	contact character varying(256) NOT NULL,
	phone character varying(256) NOT NULL,
	email character varying(256) NOT NULL,
	"partnerId" character varying(256) NOT NULL,
	"contractNum" character varying(256) NOT NULL,
	"siteName" character varying(256) NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "ServiceProvider_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."ServiceProviderHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."ServiceProviderHistory" CASCADE;
CREATE TABLE public."ServiceProviderHistory"(
	id serial NOT NULL,
	"spId" integer NOT NULL,
	"partnerName" character varying(256) NOT NULL,
	contact character varying(256) NOT NULL,
	phone character varying(256) NOT NULL,
	email character varying(256) NOT NULL,
	"partnerId" character varying(256) NOT NULL,
	"contractNum" character varying(256) NOT NULL,
	"siteName" character varying(256) NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "ServiceProviderHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."ConfigKV" | type: TABLE --
-- DROP TABLE IF EXISTS public."ConfigKV" CASCADE;
CREATE TABLE public."ConfigKV"(
	id serial NOT NULL,
	"configTypeId" integer NOT NULL,
	key character varying(256) NOT NULL,
	value character varying(1024) NOT NULL,
	"dataType" character varying(32) NOT NULL,
	active boolean NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "ConfigKV_pk" PRIMARY KEY (id),
	CONSTRAINT "ConfigKV_uk" UNIQUE ("configTypeId",key)

);
-- ddl-end --

-- object: public."ConfigKVHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."ConfigKVHistory" CASCADE;
CREATE TABLE public."ConfigKVHistory"(
	id serial NOT NULL,
	"configKVId" integer NOT NULL,
	"configTypeId" integer NOT NULL,
	key character varying(256) NOT NULL,
	value character varying(1024) NOT NULL,
	"dataType" character varying(32) NOT NULL,
	active boolean NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "ConfigKVHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."ConfigType" | type: TABLE --
-- DROP TABLE IF EXISTS public."ConfigType" CASCADE;
CREATE TABLE public."ConfigType"(
	id integer NOT NULL,
	name character varying(128) NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "ConfigType_pk" PRIMARY KEY (id),
	CONSTRAINT "ConfigType_uk" UNIQUE (name)

);
-- ddl-end --

-- object: public."Xfer" | type: TABLE --
-- DROP TABLE IF EXISTS public."Xfer" CASCADE;
CREATE TABLE public."Xfer"(
	id bigint NOT NULL,
	"changeStartTime" timestamp with time zone NOT NULL,
	"changeEndTime" timestamp with time zone NOT NULL,
	"scanStartTime" timestamp with time zone,
	"scanEndTime" timestamp with time zone,
	"xferStartTime" timestamp with time zone,
	"xferEndTime" timestamp with time zone,
	"tableCnt" integer,
	"changedTableCnt" integer,
	"changedRowCnt" bigint,
	status character varying(64),
	"changedByUserId" integer NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "Xfer_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."XferTableInfo" | type: TABLE --
-- DROP TABLE IF EXISTS public."XferTableInfo" CASCADE;
CREATE TABLE public."XferTableInfo"(
	id bigserial NOT NULL,
	"xferId" bigint NOT NULL,
	"tableName" character varying(64) NOT NULL,
	"changedRowCnt" bigint NOT NULL,
	status character varying(64) NOT NULL,
	"changedByUserId" integer NOT NULL,
	"changeTime" timestamp with time zone NOT NULL
);
-- ddl-end --

-- object: public."VcProperty" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcProperty" CASCADE;
CREATE TABLE public."VcProperty"(
	id integer NOT NULL,
	"propertyVersion" smallint NOT NULL,
	active boolean NOT NULL,
	category character varying(64) NOT NULL,
	"objType" character varying(64) NOT NULL,
	"pathSet" character varying(128) NOT NULL,
	"dataType" character varying(64) NOT NULL,
	"systemUser" boolean,
	"retrievePath" character varying(256),
	"tableName" character varying(128) NOT NULL,
	"columnName" character varying(128) NOT NULL,
	anonymize boolean NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcProperty_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcPropertyHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcPropertyHistory" CASCADE;
CREATE TABLE public."VcPropertyHistory"(
	id integer NOT NULL,
	"propertyId" integer NOT NULL,
	"propertyVersion" smallint NOT NULL,
	active boolean NOT NULL,
	category character varying(64) NOT NULL,
	"objType" character varying(64) NOT NULL,
	"pathSet" character varying(128) NOT NULL,
	"dataType" character varying(64) NOT NULL,
	"systemUser" boolean,
	"retrievePath" character varying(256),
	"tableName" character varying(128) NOT NULL,
	"columnName" character varying(128) NOT NULL,
	anonymize boolean NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcPropertyHisotry_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcCollection" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcCollection" CASCADE;
CREATE TABLE public."VcCollection"(
	id bigserial NOT NULL,
	"vcId" integer NOT NULL,
	"propertyVersion" smallint NOT NULL,
	"collectionParentId" bigint,
	"collectionType" character varying(32) NOT NULL,
	"wfuVersion" varchar(128),
	"itemCount" integer NOT NULL,
	"collectionCount" integer NOT NULL,
	"collectionStartTime" timestamp with time zone NOT NULL,
	"collectionEndTime" timestamp with time zone,
	"recordStartTime" timestamp with time zone,
	"recordEndTime" timestamp with time zone,
	status character varying(64) NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcCollection_pk" PRIMARY KEY (id),
	CONSTRAINT "VcCollection_uk" UNIQUE (id,"vcId","propertyVersion")

);
-- ddl-end --

-- object: public."VcUpdates" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcUpdates" CASCADE;
CREATE TABLE public."VcUpdates"(
	id bigserial NOT NULL,
	"vcCollectionId" bigint NOT NULL,
	"vcId" integer NOT NULL,
	"propertyVersion" smallint NOT NULL,
	moref character varying(256) NOT NULL,
	"vcPropertyId" integer NOT NULL,
	value character varying(1024),
	"eventTime" timestamp with time zone,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcUpdates_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcBaseVm" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcBaseVm" CASCADE;
CREATE TABLE public."VcBaseVm"(
	id bigserial NOT NULL,
	"vcCollectionId" bigint NOT NULL,
	"vcId" integer NOT NULL,
	"propertyVersion" smallint NOT NULL,
	moref character varying(256) NOT NULL,
	"cpuReservation" integer,
	"guestFullName" character varying(256),
	"instanceUuid" character varying,
	"memoryReservation" integer,
	"memorySizeMB" integer,
	name character varying(256),
	"numCpu" smallint,
	uuid character varying(256),
	"bootTime" timestamp with time zone,
	host character varying(256),
	"powerState" character varying(32),
	"suspendInterval" integer,
	"suspendTime" timestamp with time zone,
	"hostName" character varying(256),
	"ipAddress" character varying(256),
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcBaseVm_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcBaseHost" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcBaseHost" CASCADE;
CREATE TABLE public."VcBaseHost"(
	id bigserial NOT NULL,
	"vcCollectionId" bigint NOT NULL,
	"vcId" integer NOT NULL,
	"propertyVersion" smallint NOT NULL,
	moref character varying(256) NOT NULL,
	hz bigint,
	"numCpuCores" integer,
	"numCpuPackages" integer,
	"numCpuThreads" integer,
	"memorySize" bigint,
	uuid character varying(256),
	"fullName" character varying(256),
	"instanceUuid" character varying(256),
	"licenseProductName" character varying(256),
	"licenseProductVersion" character varying(256),
	name character varying(256),
	"productLineId" character varying(256),
	version character varying,
	"bootTime" timestamp with time zone,
	"powerState" character varying(32),
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcBaseHost_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcBaseLicense" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcBaseLicense" CASCADE;
CREATE TABLE public."VcBaseLicense"(
	id serial NOT NULL,
	"vcCollectionId" bigint NOT NULL,
	"vcId" integer NOT NULL,
	"propertyVersion" smallint NOT NULL,
	"editionKey" character varying(256),
	"licenseKey" character varying(256),
	name character varying(256),
	total smallint,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcLicenseBase_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcVm" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcVm" CASCADE;
CREATE TABLE public."VcVm"(
	id serial NOT NULL,
	"vcId" integer NOT NULL,
	moref character varying(256) NOT NULL,
	"guestFullName" character varying,
	"hostName" character varying(256),
	"instanceUuid" character varying,
	name character varying(256) NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone,
	CONSTRAINT "VcVm_pk" PRIMARY KEY (id),
	CONSTRAINT "VcVm_uk" UNIQUE (moref)

);
-- ddl-end --

-- object: public."VcHost" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcHost" CASCADE;
CREATE TABLE public."VcHost"(
	id serial NOT NULL,
	"vcId" integer NOT NULL,
	moref character varying(256) NOT NULL,
	"fullName" integer,
	name character varying(256),
	uuid character varying,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcHost_pk" PRIMARY KEY (id),
	CONSTRAINT "VcHost_uk" UNIQUE (moref)

);
-- ddl-end --

-- object: public."VcCollectInterval" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcCollectInterval" CASCADE;
CREATE TABLE public."VcCollectInterval"(
	id bigserial NOT NULL,
	"vcId" integer NOT NULL,
	"stateOn" boolean NOT NULL,
	"vcEndPoint" character varying(256) NOT NULL,
	"vcInstanceUuid" character varying(256) NOT NULL,
	"vcFullName" character varying(256) NOT NULL,
	"vcVersion" character varying(256) NOT NULL,
	"vcLicenseKey" character varying(256) NOT NULL,
	"vcAboutInfoFullName" character varying(256) NOT NULL,
	"vcLmLicensInfoName" character varying(256) NOT NULL,
	"startTime" timestamp with time zone NOT NULL,
	"endTime" timestamp with time zone NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcCollectInterval_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcMeterInterval" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcMeterInterval" CASCADE;
CREATE TABLE public."VcMeterInterval"(
	id serial NOT NULL,
	version integer NOT NULL,
	"vcId" integer NOT NULL,
	"stateOn" boolean,
	"startTime" timestamp with time zone,
	"endTime" timestamp with time zone,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcMeterInterval_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcState" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcState" CASCADE;
CREATE TABLE public."VcState"(
	"vcId" integer NOT NULL,
	"wfuOn" boolean NOT NULL,
	"wfuUpdateTime" timestamp with time zone NOT NULL,
	"containerViewOn" boolean NOT NULL,
	"containerViewUpdateTime" timestamp with time zone NOT NULL,
	"newConnectionOn" boolean NOT NULL,
	"newConnectionUpdateTime" timestamp with time zone NOT NULL,
	"errorCode" smallint,
	"retryPause" boolean,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL
);
-- ddl-end --

-- object: public."VcStateHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcStateHistory" CASCADE;
CREATE TABLE public."VcStateHistory"(
	id bigserial NOT NULL,
	"vcId" integer NOT NULL,
	"wfuOn" boolean NOT NULL,
	"wfuUpdateTime" timestamp with time zone NOT NULL,
	"containerViewOn" boolean NOT NULL,
	"containerViewUpdateTime" timestamp with time zone,
	"newConnectionOn" boolean NOT NULL,
	"newConnectionUpdateTime" timestamp NOT NULL,
	"errorCode" smallint,
	"retryPause" smallint,
	message character varying(2048),
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcStateHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcPerfStatsHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcPerfStatsHistory" CASCADE;
CREATE TABLE public."VcPerfStatsHistory"(
	id bigserial NOT NULL,
	"vcId" integer NOT NULL,
	"minSecond" real,
	"maxSecond" real,
	"avgSecond" real,
	"durationSecond" real,
	cnt integer,
	"statsType" varchar,
	"changeTime" timestamp with time zone,
	CONSTRAINT "VcPerfStatsHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."ErrorCode" | type: TABLE --
-- DROP TABLE IF EXISTS public."ErrorCode" CASCADE;
CREATE TABLE public."ErrorCode"(
	id smallint NOT NULL,
	"errorCode" smallint NOT NULL,
	"retryPause" boolean NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "ErrorCode_pk" PRIMARY KEY (id),
	CONSTRAINT "ErrorCode_uk" UNIQUE ("errorCode")

);
-- ddl-end --

-- object: public."VcCollectionRaw" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcCollectionRaw" CASCADE;
CREATE TABLE public."VcCollectionRaw"(
	id serial NOT NULL,
	"vcCollectionId" bigint NOT NULL,
	"vcId" integer NOT NULL,
	"propertyVersion" smallint NOT NULL,
	raw jsonb,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcCollectionRaw_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: public."VcLicenseCategory" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcLicenseCategory" CASCADE;
CREATE TABLE public."VcLicenseCategory"(
	id serial NOT NULL,
	"editionKey" character varying(256) NOT NULL,
	"licenseKey" character varying(256) NOT NULL,
	name character varying(256) NOT NULL,
	total smallint NOT NULL,
	"licenseCategoryId" smallint NOT NULL,
	"licenseCategoryVersion" integer NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcLicenseCategory_pk" PRIMARY KEY (id),
	CONSTRAINT "VcLicenseCategory_uk" UNIQUE ("licenseKey")

);
-- ddl-end --

-- object: public."VcLicenseCategoryHistory" | type: TABLE --
-- DROP TABLE IF EXISTS public."VcLicenseCategoryHistory" CASCADE;
CREATE TABLE public."VcLicenseCategoryHistory"(
	id bigserial NOT NULL,
	"vcLicenseId" integer NOT NULL,
	"editionKey" character varying(256) NOT NULL,
	"licenseKey" character varying(256) NOT NULL,
	name character varying(256) NOT NULL,
	total smallint NOT NULL,
	"licenseCategoryId" smallint NOT NULL,
	"licenseCategoryVersion" integer NOT NULL,
	"changedByUserId" smallint NOT NULL,
	"changeTime" timestamp with time zone NOT NULL,
	CONSTRAINT "VcLicenseCategoryHistory_pk" PRIMARY KEY (id)

);
-- ddl-end --

-- object: "PscHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."PscHistory" DROP CONSTRAINT IF EXISTS "PscHistory_fk" CASCADE;
ALTER TABLE public."PscHistory" ADD CONSTRAINT "PscHistory_fk" FOREIGN KEY ("pscId")
REFERENCES public."Psc" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "Vc_fk" | type: CONSTRAINT --
-- ALTER TABLE public."Vc" DROP CONSTRAINT IF EXISTS "Vc_fk" CASCADE;
ALTER TABLE public."Vc" ADD CONSTRAINT "Vc_fk" FOREIGN KEY ("pscId")
REFERENCES public."Psc" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcHistory" DROP CONSTRAINT IF EXISTS "VcHistory_fk" CASCADE;
ALTER TABLE public."VcHistory" ADD CONSTRAINT "VcHistory_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcRegistrationHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcRegistrationHistory" DROP CONSTRAINT IF EXISTS "VcRegistrationHistory_fk" CASCADE;
ALTER TABLE public."VcRegistrationHistory" ADD CONSTRAINT "VcRegistrationHistory_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "UserHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."UserHistory" DROP CONSTRAINT IF EXISTS "UserHistory_fk" CASCADE;
ALTER TABLE public."UserHistory" ADD CONSTRAINT "UserHistory_fk" FOREIGN KEY ("userId")
REFERENCES public."User" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "UserLoginHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."UserLoginHistory" DROP CONSTRAINT IF EXISTS "UserLoginHistory_fk" CASCADE;
ALTER TABLE public."UserLoginHistory" ADD CONSTRAINT "UserLoginHistory_fk" FOREIGN KEY ("userId")
REFERENCES public."User" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "UserPriv_fk" | type: CONSTRAINT --
-- ALTER TABLE public."UserPriv" DROP CONSTRAINT IF EXISTS "UserPriv_fk" CASCADE;
ALTER TABLE public."UserPriv" ADD CONSTRAINT "UserPriv_fk" FOREIGN KEY ("roleId")
REFERENCES public."Role" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "UserPriv_fk01" | type: CONSTRAINT --
-- ALTER TABLE public."UserPriv" DROP CONSTRAINT IF EXISTS "UserPriv_fk01" CASCADE;
ALTER TABLE public."UserPriv" ADD CONSTRAINT "UserPriv_fk01" FOREIGN KEY ("userId")
REFERENCES public."User" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "UserPrivHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."UserPrivHistory" DROP CONSTRAINT IF EXISTS "UserPrivHistory_fk" CASCADE;
ALTER TABLE public."UserPrivHistory" ADD CONSTRAINT "UserPrivHistory_fk" FOREIGN KEY ("userPrivId")
REFERENCES public."UserPriv" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "ServiceProviderHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."ServiceProviderHistory" DROP CONSTRAINT IF EXISTS "ServiceProviderHistory_fk" CASCADE;
ALTER TABLE public."ServiceProviderHistory" ADD CONSTRAINT "ServiceProviderHistory_fk" FOREIGN KEY ("spId")
REFERENCES public."ServiceProvider" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "ConfigKV_fk" | type: CONSTRAINT --
-- ALTER TABLE public."ConfigKV" DROP CONSTRAINT IF EXISTS "ConfigKV_fk" CASCADE;
ALTER TABLE public."ConfigKV" ADD CONSTRAINT "ConfigKV_fk" FOREIGN KEY ("configTypeId")
REFERENCES public."ConfigType" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "ConfigKVHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."ConfigKVHistory" DROP CONSTRAINT IF EXISTS "ConfigKVHistory_fk" CASCADE;
ALTER TABLE public."ConfigKVHistory" ADD CONSTRAINT "ConfigKVHistory_fk" FOREIGN KEY ("configKVId")
REFERENCES public."ConfigKV" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "XferTableInfo_fk" | type: CONSTRAINT --
-- ALTER TABLE public."XferTableInfo" DROP CONSTRAINT IF EXISTS "XferTableInfo_fk" CASCADE;
ALTER TABLE public."XferTableInfo" ADD CONSTRAINT "XferTableInfo_fk" FOREIGN KEY ("xferId")
REFERENCES public."Xfer" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcCollectionVcCollection_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcCollection" DROP CONSTRAINT IF EXISTS "VcCollectionVcCollection_fk" CASCADE;
ALTER TABLE public."VcCollection" ADD CONSTRAINT "VcCollectionVcCollection_fk" FOREIGN KEY ("collectionParentId")
REFERENCES public."VcCollection" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcVcCollection_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcCollection" DROP CONSTRAINT IF EXISTS "VcVcCollection_fk" CASCADE;
ALTER TABLE public."VcCollection" ADD CONSTRAINT "VcVcCollection_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcCollectionVcUpdates_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcUpdates" DROP CONSTRAINT IF EXISTS "VcCollectionVcUpdates_fk" CASCADE;
ALTER TABLE public."VcUpdates" ADD CONSTRAINT "VcCollectionVcUpdates_fk" FOREIGN KEY ("vcCollectionId","vcId","propertyVersion")
REFERENCES public."VcCollection" (id,"vcId","propertyVersion") MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcPropertyVcUpdates_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcUpdates" DROP CONSTRAINT IF EXISTS "VcPropertyVcUpdates_fk" CASCADE;
ALTER TABLE public."VcUpdates" ADD CONSTRAINT "VcPropertyVcUpdates_fk" FOREIGN KEY ("vcPropertyId")
REFERENCES public."VcProperty" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcCollectionVcBaseVm_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcBaseVm" DROP CONSTRAINT IF EXISTS "VcCollectionVcBaseVm_fk" CASCADE;
ALTER TABLE public."VcBaseVm" ADD CONSTRAINT "VcCollectionVcBaseVm_fk" FOREIGN KEY ("vcCollectionId","vcId","propertyVersion")
REFERENCES public."VcCollection" (id,"vcId","propertyVersion") MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcCollectionVcBaseHost_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcBaseHost" DROP CONSTRAINT IF EXISTS "VcCollectionVcBaseHost_fk" CASCADE;
ALTER TABLE public."VcBaseHost" ADD CONSTRAINT "VcCollectionVcBaseHost_fk" FOREIGN KEY ("vcCollectionId","vcId","propertyVersion")
REFERENCES public."VcCollection" (id,"vcId","propertyVersion") MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcCollectionVcLicenseBase_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcBaseLicense" DROP CONSTRAINT IF EXISTS "VcCollectionVcLicenseBase_fk" CASCADE;
ALTER TABLE public."VcBaseLicense" ADD CONSTRAINT "VcCollectionVcLicenseBase_fk" FOREIGN KEY ("vcCollectionId","vcId","propertyVersion")
REFERENCES public."VcCollection" (id,"vcId","propertyVersion") MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcVcVm_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcVm" DROP CONSTRAINT IF EXISTS "VcVcVm_fk" CASCADE;
ALTER TABLE public."VcVm" ADD CONSTRAINT "VcVcVm_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcVcHost_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcHost" DROP CONSTRAINT IF EXISTS "VcVcHost_fk" CASCADE;
ALTER TABLE public."VcHost" ADD CONSTRAINT "VcVcHost_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcVcStateInterval_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcCollectInterval" DROP CONSTRAINT IF EXISTS "VcVcStateInterval_fk" CASCADE;
ALTER TABLE public."VcCollectInterval" ADD CONSTRAINT "VcVcStateInterval_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcVcMeteringHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcMeterInterval" DROP CONSTRAINT IF EXISTS "VcVcMeteringHistory_fk" CASCADE;
ALTER TABLE public."VcMeterInterval" ADD CONSTRAINT "VcVcMeteringHistory_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcVcState_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcState" DROP CONSTRAINT IF EXISTS "VcVcState_fk" CASCADE;
ALTER TABLE public."VcState" ADD CONSTRAINT "VcVcState_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcVcStateHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcStateHistory" DROP CONSTRAINT IF EXISTS "VcVcStateHistory_fk" CASCADE;
ALTER TABLE public."VcStateHistory" ADD CONSTRAINT "VcVcStateHistory_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcVcPerfStatsHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcPerfStatsHistory" DROP CONSTRAINT IF EXISTS "VcVcPerfStatsHistory_fk" CASCADE;
ALTER TABLE public."VcPerfStatsHistory" ADD CONSTRAINT "VcVcPerfStatsHistory_fk" FOREIGN KEY ("vcId")
REFERENCES public."Vc" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcCollectionVcCollectionRaw_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcCollectionRaw" DROP CONSTRAINT IF EXISTS "VcCollectionVcCollectionRaw_fk" CASCADE;
ALTER TABLE public."VcCollectionRaw" ADD CONSTRAINT "VcCollectionVcCollectionRaw_fk" FOREIGN KEY ("vcCollectionId","vcId","propertyVersion")
REFERENCES public."VcCollection" (id,"vcId","propertyVersion") MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: "VcLicenseCategoryVcLicenseCategoryHistory_fk" | type: CONSTRAINT --
-- ALTER TABLE public."VcLicenseCategoryHistory" DROP CONSTRAINT IF EXISTS "VcLicenseCategoryVcLicenseCategoryHistory_fk" CASCADE;
ALTER TABLE public."VcLicenseCategoryHistory" ADD CONSTRAINT "VcLicenseCategoryVcLicenseCategoryHistory_fk" FOREIGN KEY ("vcLicenseId")
REFERENCES public."VcLicenseCategory" (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


